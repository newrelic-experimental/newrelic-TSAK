package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  flow "github.com/kamildrazkiewicz/go-flow"
)


func init() {
  env.Packages["stdlib/flow"] = map[string]reflect.Value{
    "New":                  reflect.ValueOf(flow.New),
  }
  env.PackageTypes["stdlib/flow"] = map[string]reflect.Type{
    "Flow":                  reflect.TypeOf(flow.Flow{}),
  }
}
