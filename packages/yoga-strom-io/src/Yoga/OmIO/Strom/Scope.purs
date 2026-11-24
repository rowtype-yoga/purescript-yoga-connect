-- | Where a running stream keeps what it must clean up, and how it learns that it should stop.
-- |
-- | A scope has a Go `context.Context` and finalizers. Closing it cancels the context, so whatever
-- | runs in it and waits (to hand a chunk over, to sleep, for a gRPC message) stops waiting, and
-- | then runs the finalizers, last added first, each with how the scope ended: `Ended`, `Failed`
-- | or `Interrupted`.
-- |
-- | Scopes nest. A child closes with its parent, as `Interrupted`, unless it closed before.
-- | Scopes are safe to share between goroutines.
module Yoga.OmIO.Strom.Scope
  ( Scope
  , Exit(..)
  , Failure
  , FinalizerKey
  , scoped
  , scopedWithin
  , forkScope
  , forkScopeWithTimeout
  , closeScope
  , addFinalizer
  , removeFinalizer
  , interrupt
  , isInterrupted
  , deadlinePassedIn
  , contextOf
  , sleep
  , attempt
  , interrupted
  , isInterruption
  , throwIfInterrupted
  , exitAfter
  ) where

import Prelude

import Control.Monad.Error.Class (throwError)
import Control.Monad.Reader (ask)
import Data.Array as Array
import Data.Either (Either(..), either)
import Data.Maybe (Maybe(..))
import Data.Time.Duration (Milliseconds)
import Data.Traversable (for)
import Data.Tuple.Nested (type (/\), (/\))
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect.Exception (Error, error, message)
import Go.IO (IO)
import Go.IO as IO
import Go.IO.IORef (IORef, atomicModifyIORef', newIORef, readIORef, writeIORef)
import Type.Proxy (Proxy(..))
import Yoga.Go.HTTP.OmIO (OmIO, liftIO, runOmIO)
import Yoga.OmIO.Strom.Context (CancelContext, Context, background, cancel, deadlinePassed, isDone, sleepUnlessDone, withCancel, withTimeout)

newtype Scope ctx err = Scope
  { context :: Context
  , cancelContext :: CancelContext
  , finalizers :: IORef (Finalizers ctx err)
  , detachFromParent :: IORef (IO Unit)
  }

-- | How a stream fails: with one of its own errors, or with an exception.
type Failure err = Variant (exception :: Error | err)

-- | How a scope ended: its stream ran out, failed, or was stopped from outside, because whoever
-- | read it was done or its context was cancelled.
data Exit err
  = Ended
  | Failed (Failure err)
  | Interrupted

-- | While open, the key the next finalizer gets, and the finalizers in the order they came.
-- | (Data rather than records, which Iris boxes when they go through an `IORef`.)
data Finalizers ctx err
  = Open FinalizerKey (Array (Registered ctx err))
  | Closed

data Registered ctx err = Registered FinalizerKey (Exit err -> OmIO ctx err Unit)

-- | Says which finalizer to remove.
newtype FinalizerKey = FinalizerKey Int

derive instance Eq FinalizerKey

-- | Runs the action in a new scope and closes it afterwards, however the action ends.
scoped :: forall ctx err a. (Scope ctx err -> OmIO ctx err a) -> OmIO ctx err a
scoped use = scopedWithin background use

-- | `scoped`, in a scope that is interrupted once the outside context is done: a gRPC call's.
scopedWithin :: forall ctx err a. Context -> (Scope ctx err -> OmIO ctx err a) -> OmIO ctx err a
scopedWithin outside use = do
  scope <- newScope outside # liftIO
  outcome <- attempt (use scope)
  exit <- exitOf scope outcome
  closed <- attempt (closeScope exit scope)
  either throwError pure (outcome <* closed)
  where
  exitOf scope = case _ of
    Right _ -> pure Ended
    Left failure -> exitAfter scope failure

-- | `Interrupted` if the failure comes from the scope being interrupted, `Failed` otherwise.
exitAfter :: forall ctx err. Scope ctx err -> Failure err -> OmIO ctx err (Exit err)
exitAfter scope failure = do
  stopped <- isInterrupted scope
  pure if stopped || isInterruption failure then Interrupted else Failed failure

newScope :: forall ctx err. Context -> IO (Scope ctx err)
newScope parentContext = do
  context /\ cancelContext <- withCancel parentContext
  scopeWith context cancelContext

scopeWith :: forall ctx err. Context -> CancelContext -> IO (Scope ctx err)
scopeWith context cancelContext = do
  finalizers <- newIORef (Open (FinalizerKey 0) [])
  detachFromParent <- newIORef (pure unit)
  pure (Scope { context, cancelContext, finalizers, detachFromParent })

-- | A child of the scope: it closes with its parent, unless it closed before.
forkScope :: forall ctx err. Scope ctx err -> OmIO ctx err (Scope ctx err)
forkScope parent = do
  child <- newScope (contextOf parent) # liftIO
  adopt parent child

-- | A child that is interrupted by itself after the duration.
forkScopeWithTimeout :: forall ctx err. Milliseconds -> Scope ctx err -> OmIO ctx err (Scope ctx err)
forkScopeWithTimeout duration parent = do
  context /\ cancelContext <- withTimeout duration (contextOf parent) # liftIO
  child <- scopeWith context cancelContext # liftIO
  adopt parent child

adopt :: forall ctx err. Scope ctx err -> Scope ctx err -> OmIO ctx err (Scope ctx err)
adopt parent child@(Scope { detachFromParent }) = do
  key <- addFinalizer parent \_ -> closeScope Interrupted child
  writeIORef detachFromParent (removeFinalizer parent key) # liftIO
  pure child

-- | Runs the finalizer when the scope closes. If it closed already, runs it right away.
addFinalizer :: forall ctx err. Scope ctx err -> (Exit err -> OmIO ctx err Unit) -> OmIO ctx err FinalizerKey
addFinalizer (Scope { finalizers }) finalize = do
  added <- atomicModifyIORef' finalizers register # liftIO
  case added of
    Just key -> pure key
    Nothing -> finalize Interrupted $> FinalizerKey (-1)
  where
  register = case _ of
    Open key registered -> Open (next key) (Array.snoc registered (Registered key finalize)) /\ Just key
    Closed -> Closed /\ Nothing

  next (FinalizerKey key) = FinalizerKey (key + 1)

removeFinalizer :: forall ctx err. Scope ctx err -> FinalizerKey -> IO Unit
removeFinalizer (Scope { finalizers }) key = atomicModifyIORef' finalizers (\state -> withoutKey state /\ unit)
  where
  withoutKey = case _ of
    Open nextKey registered -> Open nextKey (Array.filter (not <<< hasKey) registered)
    Closed -> Closed

  hasKey (Registered registeredKey _) = registeredKey == key

-- | Cancels the scope's context and runs its finalizers, last added first, each to the end even
-- | if one fails; then fails with the first failure. Closing a closed scope does nothing.
closeScope :: forall ctx err. Exit err -> Scope ctx err -> OmIO ctx err Unit
closeScope exit (Scope { cancelContext, finalizers, detachFromParent }) = do
  state <- atomicModifyIORef' finalizers (\state -> Closed /\ state) # liftIO
  case state of
    Closed -> pure unit
    Open _ registered -> do
      cancel cancelContext # liftIO
      detach <- readIORef detachFromParent # liftIO
      detach # liftIO
      outcomes <- for (Array.reverse registered) \(Registered _ finalize) -> attempt (finalize exit)
      case Array.find isFailure outcomes of
        Just (Left failure) -> throwError failure
        _ -> pure unit
  where
  isFailure = either (const true) (const false)

-- | Cancels the scope's context without running its finalizers, so that whatever runs in it
-- | stops at its next wait and closes the scope itself.
interrupt :: forall ctx err. Scope ctx err -> IO Unit
interrupt (Scope { cancelContext }) = cancel cancelContext

isInterrupted :: forall ctx err. Scope ctx err -> OmIO ctx err Boolean
isInterrupted scope = isDone (contextOf scope) # liftIO

deadlinePassedIn :: forall ctx err. Scope ctx err -> OmIO ctx err Boolean
deadlinePassedIn scope = deadlinePassed (contextOf scope) # liftIO

-- | The context to hand to Go calls made for the scope, such as a gRPC call.
contextOf :: forall ctx err. Scope ctx err -> Context
contextOf (Scope { context }) = context

-- | Waits the duration, and fails with `interrupted` as soon as the scope is.
sleep :: forall ctx err. Scope ctx err -> Milliseconds -> OmIO ctx err Unit
sleep scope duration = do
  slept <- sleepUnlessDone (contextOf scope) duration # liftIO
  unless slept (throwError interrupted)

throwIfInterrupted :: forall ctx err. Scope ctx err -> OmIO ctx err Unit
throwIfInterrupted scope = do
  stopped <- isInterrupted scope
  when stopped (throwError interrupted)

-- | What a stream fails with when it stops because its scope was interrupted.
interrupted :: forall err. Failure err
interrupted = Variant.inj (Proxy :: Proxy "exception") interruption

interruption :: Error
interruption = error interruptionMessage

interruptionMessage :: String
interruptionMessage = "the stream was interrupted"

isInterruption :: forall err. Failure err -> Boolean
isInterruption = Variant.on (Proxy :: Proxy "exception") (\thrown -> message thrown == interruptionMessage) (const false)

-- | The action's outcome, with whatever it throws in `IO`, such as a panic in Go, as an
-- | `exception`.
attempt :: forall ctx err a. OmIO ctx err a -> OmIO ctx err (Either (Failure err) a)
attempt action = do
  context <- ask
  outcome <- IO.try (runOmIO context action) # liftIO
  pure (either (Left <<< Variant.inj (Proxy :: Proxy "exception")) identity outcome)
