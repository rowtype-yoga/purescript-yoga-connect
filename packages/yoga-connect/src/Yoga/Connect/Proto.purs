-- | The `.proto` of a service type, for grpcurl and for other languages' code generators.
-- |
-- | ```purescript
-- | log (printProto @Ticker)
-- | ```
-- |
-- | A method's messages are `WatchRequest` and `WatchResponse`, unless `Named` names them. A
-- | nested message or enum is named after its field, or by `Named`, and goes at the top of the
-- | file when every one of that name is the same, inside the message that has it when they
-- | differ. An error that carries a value has a message of its own, `ColumnFull`, for the detail
-- | it goes out with. Field names are in snake case, so that their JSON names are the record's labels
-- | again: `averagePrice` is `average_price = 3`, which JSON calls `averagePrice`.
module Yoga.Connect.Proto
  ( printProto
  , printService
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String (Pattern(..))
import Data.String as String
import Data.String.CodeUnits as CodeUnits
import Data.Symbol (class IsSymbol, reflectSymbol)
import Data.Tuple.Nested (type (/\), (/\))
import Prim.RowList (class RowToList)
import Type.Proxy (Proxy(..))
import Yoga.Connect.Protobuf (Cardinality(..), FieldShape, Member(..), MessageShape, ValueKind(..))
import Yoga.Connect.Schema (class ServiceMethods, class ServiceType, MethodShape, methodNameOf, methodShapes)

printProto
  :: forall @service name methods list connect
   . ServiceType service name methods
  => IsSymbol name
  => RowToList methods list
  => ServiceMethods list connect
  => String
printProto = printService (reflectSymbol (Proxy :: Proxy name)) (methodShapes (Proxy :: Proxy list))

-- | The `.proto` of the service with this full name and these methods.
printService :: String -> Array { name :: String, schema :: MethodShape } -> String
printService fullName methods = String.joinWith "\n" (Array.concat [ header, [ service ], methodMessages, topLevel ])
  where
  { package, serviceName } = splitName fullName
  header = [ "syntax = \"proto3\";\n" ] <> (package # foldMaybe \name -> [ "package " <> name <> ";\n" ])
  service = "service " <> serviceName <> " {\n" <> foldLines (methods <#> rpc) <> "}\n"
  rpc { name, schema } =
    "  rpc " <> name <> "(" <> streamed schema.clientStreams <> requestOf name schema <> ") returns ("
      <> streamed schema.serverStreams
      <> responseOf name schema
      <> ");\n"
  streamed streams = if streams then "stream " else ""
  requestOf name schema = schema.requestName # fromMaybe (name <> "Request")
  responseOf name schema = schema.responseName # fromMaybe (name <> "Response")

  ownMessages = Array.nubBy (comparing fst') (methods >>= \{ name, schema } -> [ requestOf name schema /\ schema.request, responseOf name schema /\ schema.response ] <> (schema.errors <#> \error -> error.name /\ error.shape))
  methodMessages = ownMessages <#> \(name /\ shape) -> printMessage hoisted "" name shape

  definitions = methods >>= \{ schema } -> nestedIn schema.request <> nestedIn schema.response <> (schema.errors >>= \error -> nestedIn error.shape)
  hoisted = Map.fromFoldableWith (\left right -> if left == right then left else Conflicting) (definitions <#> map Single')
    # Map.mapMaybe case _ of
        Single' definition -> Just definition
        Conflicting -> Nothing
  topLevel = Map.toUnfoldable hoisted # Array.filter (\(name /\ _) -> not (Array.elem name (ownMessages <#> fst'))) <#> \(name /\ definition) -> printDefinition hoisted "" name definition
  fst' (name /\ _) = name
  foldMaybe f = case _ of
    Just value -> f value
    Nothing -> []

-- | A message or an enum, under its name.
data Definition
  = MessageDefinition MessageShape
  | EnumDefinition (Array String)

derive instance Eq Definition

-- | One definition of a name, or several that differ.
data Hoisting = Single' Definition | Conflicting

derive instance Eq Hoisting

-- | Every message and enum the fields of a message use, all the way down.
nestedIn :: MessageShape -> Array (String /\ Definition)
nestedIn members = members >>= memberFields >>= \field -> case definitionOf field of
  Nothing -> []
  Just (name /\ definition@(MessageDefinition inner)) -> [ name /\ definition ] <> nestedIn inner
  Just found -> [ found ]

definitionOf :: FieldShape -> Maybe (String /\ Definition)
definitionOf field = case field.kind of
  Scalar _ -> Nothing
  Nested _ shape -> Just (typeName field /\ MessageDefinition shape)
  Enumeration _ constructors -> Just (typeName field /\ EnumDefinition constructors)

printDefinition :: Map String Definition -> String -> String -> Definition -> String
printDefinition hoisted indent name = case _ of
  MessageDefinition shape -> printMessage hoisted indent name shape
  EnumDefinition constructors -> printEnum indent name constructors

printMessage :: Map String Definition -> String -> String -> MessageShape -> String
printMessage hoisted indent name members =
  indent <> "message " <> name <> " {\n" <> foldLines (nested <> fields) <> indent <> "}\n"
  where
  inner = indent <> "  "
  nested = members >>= memberFields
    # Array.mapMaybe definitionOf
    # Array.filter (\(typeOf /\ _) -> not (Map.member typeOf hoisted))
    # Array.nubBy (comparing \(typeOf /\ _) -> typeOf)
    <#> \(typeOf /\ definition) -> printDefinition hoisted inner typeOf definition
  fields = Array.sortWith firstNumber members <#> printMember inner
  firstNumber = case _ of
    Single field -> field.number
    Choice _ alternatives -> Array.foldl (\least field -> min least field.number) top alternatives

memberFields :: Member -> Array FieldShape
memberFields = case _ of
  Single field -> [ field ]
  Choice _ alternatives -> alternatives

printMember :: String -> Member -> String
printMember indent = case _ of
  Single field -> printField indent field
  Choice name alternatives ->
    indent <> "oneof " <> snakeCase name <> " {\n" <> foldLines (Array.sortWith _.number alternatives <#> printField (indent <> "  ")) <> indent <> "}\n"

printField :: String -> FieldShape -> String
printField indent field = indent <> cardinality <> typeName field <> " " <> snakeCase field.name <> " = " <> show field.number <> ";\n"
  where
  cardinality = case field.cardinality of
    Once -> ""
    Repeated -> "repeated "
    Optional -> "optional "

printEnum :: String -> String -> Array String -> String
printEnum indent name constructors =
  indent <> "enum " <> name <> " {\n" <> foldLines (Array.mapWithIndex value constructors) <> indent <> "}\n"
  where
  value index constructor = indent <> "  " <> screamingSnakeCase name <> "_" <> screamingSnakeCase constructor <> " = " <> show index <> ";\n"

-- | A scalar's own name, a nested type's `Named` name or else its field's name.
typeName :: FieldShape -> String
typeName field = case field.kind of
  Scalar name -> name
  Nested (Just name) _ -> name
  Enumeration (Just name) _ -> name
  _ -> methodNameOf field.name

splitName :: String -> { package :: Maybe String, serviceName :: String }
splitName fullName = case String.lastIndexOf (Pattern ".") fullName of
  Just dot -> { package: Just (CodeUnits.take dot fullName), serviceName: CodeUnits.drop (dot + 1) fullName }
  Nothing -> { package: Nothing, serviceName: fullName }

-- | `averagePrice` is `average_price`.
snakeCase :: String -> String
snakeCase label = CodeUnits.fromCharArray (CodeUnits.toCharArray label >>= underscored)
  where
  underscored character
    | isUpper character = [ '_' ] <> CodeUnits.toCharArray (String.toLower (CodeUnits.singleton character))
    | otherwise = [ character ]
  isUpper character = character >= 'A' && character <= 'Z'

screamingSnakeCase :: String -> String
screamingSnakeCase name = String.toUpper (snakeCase (String.toLower (CodeUnits.take 1 name) <> CodeUnits.drop 1 name))

foldLines :: Array String -> String
foldLines = String.joinWith ""
