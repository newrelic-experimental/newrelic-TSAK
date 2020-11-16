package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/Jeffail/gabs"
)

func JSONParsing(data []byte) *gabs.Container {
  c, err := gabs.ParseJSON(data)
  if err != nil {
    return gabs.New()
  }
  return c
}

func init() {
  env.Packages["djson"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(gabs.New),
    "Parse":              reflect.ValueOf(JSONParsing),
  }
  env.PackageTypes["djson"] = map[string]reflect.Type{
    "Container":          reflect.TypeOf(gabs.Container{}),
  }
}
