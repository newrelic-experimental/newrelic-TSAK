package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/yourbasic/bloom"
)

func init() {
  env.Packages["filter/bloom"] = map[string]reflect.Value{
    "New":             reflect.ValueOf(bloom.New),
  }
  env.PackageTypes["filter/bloom"] = map[string]reflect.Type{
    "Filter":          reflect.TypeOf(bloom.Filter{}),
  }
}
