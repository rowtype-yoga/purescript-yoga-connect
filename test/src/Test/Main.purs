module Test.Main where

import Prelude

import Effect (Effect)
import Effect.Aff (launchAff_)
import Test.Spec.Reporter.Console (consoleReporter)
import Test.Spec.Runner (defaultConfig, runSpec')
import Test.Yoga.Connect.CallSpec as CallSpec
import Test.Yoga.Connect.ServiceSpec as ServiceSpec
import Test.Yoga.Connect.StatusSpec as StatusSpec
import Test.Yoga.Connect.StromSpec as StromSpec

main :: Effect Unit
main = launchAff_ do
  runSpec' defaultConfig [ consoleReporter ] do
    StatusSpec.spec
    ServiceSpec.spec
    StromSpec.spec
    CallSpec.spec
