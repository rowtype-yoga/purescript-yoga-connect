module Test.Yoga.OmIO.Main where

import Prelude

import Effect (Effect)
import Effect.Aff (launchAff_)
import Test.Spec.Reporter.Console (consoleReporter)
import Test.Spec.Runner (defaultConfig, runSpec')
import Test.Yoga.OmIO.StromSpec as StromSpec

main :: Effect Unit
main = launchAff_ do
  runSpec' defaultConfig [ consoleReporter ] StromSpec.spec
