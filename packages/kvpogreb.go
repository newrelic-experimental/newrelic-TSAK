package packages

import (
  "reflect"
  "github.com/akrylysov/pogreb"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["kv/pogreb"] = map[string]reflect.Value{
    "Open":               reflect.ValueOf(pogreb.Open),
  }
  env.PackageTypes["kv/pogreb"] = map[string]reflect.Type{
    "DB":                 reflect.TypeOf(pogreb.DB{}),
    "ItemIterator":       reflect.TypeOf(pogreb.ItemIterator{}),
  }
}
