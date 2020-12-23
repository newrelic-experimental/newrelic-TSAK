package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/piquette/finance-go"
  "github.com/piquette/finance-go/forex"
)

func init() {
  env.Packages["nostdlib/finance/forex"] = map[string]reflect.Value{
    "Forex":             reflect.ValueOf(forex.Get),
  }
  env.PackageTypes["nostdlib/finance/forex"] = map[string]reflect.Type{
    "ForexPair":         reflect.TypeOf(finance.ForexPair{}),
  }
}
