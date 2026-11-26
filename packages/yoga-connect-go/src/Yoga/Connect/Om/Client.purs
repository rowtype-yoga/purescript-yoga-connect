-- | A client of a service type: a record with a function per method.
-- |
-- | ```purescript
-- | let ticker = client @Ticker connection
-- | ticker.watch { symbols: [ "ACME" ] } # Strom.take 3 # Strom.forEach (log <<< show)
-- | ```
-- |
-- | | Method | Client function |
-- | | --- | --- |
-- | | `Unary req res err` | `req -> OmIO ctx failing res` |
-- | | `ServerStream req res err` | `req -> Strom ctx failing res` |
-- | | `ClientStream req res err` | `Strom ctx failing req -> OmIO ctx failing res` |
-- | | `BidiStream req res err` | `Strom ctx failing req -> Strom ctx failing res` |
-- |
-- | `failing` is the caller's error row, `(connectError :: ConnectError | …)`, and it has the
-- | labels of every method's errors: a call the server rejects with `reject @"notFound" why`
-- | fails with `notFound` and `why`. Any other failure is a `connectError` with its code.
module Yoga.Connect.Om.Client
  ( client
  , clientWithin
  , class TypedClients
  , typedClients
  , class TypedClient
  , typedClient
  , class Retyped
  , retyped
  ) where

import Prelude

import Control.Monad.Error.Class (catchError, throwError)
import Data.Maybe (Maybe(..), maybe)
import Data.Symbol (class IsSymbol, reflectSymbol)
import Data.Variant (Variant)
import Data.Variant as Variant
import Effect.Exception (Error)
import Prim.Row as Row
import Prim.RowList (class RowToList, RowList)
import Prim.RowList as RowList
import Record.Builder (Builder)
import Record.Builder as Builder
import Type.Proxy (Proxy(..))
import Yoga.Connect.Client (class ClientMethod, Connection, clientMethod)
import Yoga.Connect.Protobuf (class MethodMessage, methodCodec)
import Yoga.Connect.Schema (class ServiceType, BidiStream, ClientStream, ServerStream, Unary, procedureOf)
import Yoga.Connect.Service as Connect
import Yoga.Connect.Status (class RejectionCode, Code, ConnectError(..), rejectionCode)
import Yoga.Go.HTTP.OmIO (OmIO)
import Yoga.OmIO.Strom (Strom)
import Yoga.OmIO.Strom as Strom
import Yoga.OmIO.Strom.Context (Context, background)

-- | A function for each method of the service, whose calls have no deadline.
client
  :: forall @service name methods list ctx err clients
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => TypedClients list ctx err clients
  => Connection
  -> Record clients
client connection = Builder.buildFromScratch (typedClients (Proxy :: Proxy list) (Proxy :: Proxy ctx) (Proxy :: Proxy err) background (reflectSymbol (Proxy :: Proxy name)) connection)

-- | `client`, whose unary and client-streaming calls are made within the context: its deadline
-- | is theirs. A streamed response is made within the scope it's read in.
clientWithin
  :: forall @service name methods list ctx err clients
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => TypedClients list ctx err clients
  => Context
  -> Connection
  -> Record clients
clientWithin context connection = Builder.buildFromScratch (typedClients (Proxy :: Proxy list) (Proxy :: Proxy ctx) (Proxy :: Proxy err) context (reflectSymbol (Proxy :: Proxy name)) connection)

type Failing err = (connectError :: ConnectError | err)

class TypedClients :: RowList Type -> Type -> Row Type -> Row Type -> Constraint
class TypedClients list ctx err clients | list ctx err -> clients where
  typedClients :: Proxy list -> Proxy ctx -> Proxy err -> Context -> String -> Connection -> Builder {} (Record clients)

instance TypedClients RowList.Nil ctx err () where
  typedClients _ _ _ _ _ _ = identity

instance
  ( IsSymbol label
  , TypedClient method ctx err function
  , TypedClients tail ctx err otherClients
  , Row.Cons label function otherClients clients
  , Row.Lacks label otherClients
  ) =>
  TypedClients (RowList.Cons label method tail) ctx err clients where
  typedClients _ ctx errors context serviceName connection =
    Builder.insert (Proxy :: Proxy label) (typedClient (Proxy :: Proxy method) ctx errors context (procedureOf serviceName (reflectSymbol (Proxy :: Proxy label))) connection)
      <<< typedClients (Proxy :: Proxy tail) ctx errors context serviceName connection

-- | The function that calls one method, by its kind.
class TypedClient :: Type -> Type -> Row Type -> Type -> Constraint
class TypedClient method ctx err function | method ctx err -> function where
  typedClient :: Proxy method -> Proxy ctx -> Proxy err -> Context -> String -> Connection -> function

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , ClientMethod (Connect.Unary plainRequest plainResponse) ctx err (plainRequest -> OmIO ctx (Failing err) plainResponse)
  , RowToList errors labels
  , Retyped labels errors
  , Row.Union errors others (exception :: Error | Failing err)
  ) =>
  TypedClient (Unary request response errors) ctx err (plainRequest -> OmIO ctx (Failing err) plainResponse) where
  typedClient _ ctx err context procedure connection request =
    catchError (call request) (throwError <<< typedFailure (Proxy :: Proxy labels))
    where
    call = clientMethod ctx err context (Connect.unary procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) connection

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , ClientMethod (Connect.ServerStream plainRequest plainResponse) ctx err (plainRequest -> Strom ctx (Failing err) plainResponse)
  , RowToList errors labels
  , Retyped labels errors
  , Row.Union errors others (exception :: Error | Failing err)
  ) =>
  TypedClient (ServerStream request response errors) ctx err (plainRequest -> Strom ctx (Failing err) plainResponse) where
  typedClient _ ctx err context procedure connection request =
    call request # Strom.mapError (typedFailure (Proxy :: Proxy labels))
    where
    call = clientMethod ctx err context (Connect.serverStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) connection

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , ClientMethod (Connect.ClientStream plainRequest plainResponse) ctx err (Strom ctx (Failing err) plainRequest -> OmIO ctx (Failing err) plainResponse)
  , RowToList errors labels
  , Retyped labels errors
  , Row.Union errors others (exception :: Error | Failing err)
  ) =>
  TypedClient (ClientStream request response errors) ctx err (Strom ctx (Failing err) plainRequest -> OmIO ctx (Failing err) plainResponse) where
  typedClient _ ctx err context procedure connection requests =
    catchError (call requests) (throwError <<< typedFailure (Proxy :: Proxy labels))
    where
    call = clientMethod ctx err context (Connect.clientStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) connection

instance
  ( MethodMessage request plainRequest
  , MethodMessage response plainResponse
  , ClientMethod (Connect.BidiStream plainRequest plainResponse) ctx err (Strom ctx (Failing err) plainRequest -> Strom ctx (Failing err) plainResponse)
  , RowToList errors labels
  , Retyped labels errors
  , Row.Union errors others (exception :: Error | Failing err)
  ) =>
  TypedClient (BidiStream request response errors) ctx err (Strom ctx (Failing err) plainRequest -> Strom ctx (Failing err) plainResponse) where
  typedClient _ ctx err context procedure connection requests =
    call requests # Strom.mapError (typedFailure (Proxy :: Proxy labels))
    where
    call = clientMethod ctx err context (Connect.bidiStream procedure (methodCodec (Proxy :: Proxy request)) (methodCodec (Proxy :: Proxy response))) connection

-- | A `connectError` whose code is one of the method's errors, as that error.
typedFailure
  :: forall labels errors others err
   . Retyped labels errors
  => Row.Union errors others (exception :: Error | Failing err)
  => Proxy labels
  -> Variant (exception :: Error | Failing err)
  -> Variant (exception :: Error | Failing err)
typedFailure labels failure = failure # Variant.on (Proxy :: Proxy "connectError") asDeclared (const failure)
  where
  asDeclared (ConnectError { code, message }) = maybe failure Variant.expand (retyped labels code message)

-- | The method's error with this code, if it has one.
class Retyped :: RowList Type -> Row Type -> Constraint
class Retyped labels errors | labels -> errors where
  retyped :: Proxy labels -> Code -> String -> Maybe (Variant errors)

instance Retyped RowList.Nil () where
  retyped _ _ _ = Nothing

instance
  ( IsSymbol label
  , RejectionCode label
  , Retyped tail others
  , Row.Cons label String others errors
  , Row.Cons label String () single
  , Row.Union others single errors
  ) =>
  Retyped (RowList.Cons label String tail) errors where
  retyped _ code message
    | code == rejectionCode @label = Just (Variant.inj (Proxy :: Proxy label) message)
    | otherwise = retyped (Proxy :: Proxy tail) code message <#> Variant.expand
