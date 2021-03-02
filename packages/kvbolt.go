package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/boltdb/bolt"
)

func init() {
  env.Packages["kv/bolt"] = map[string]reflect.Value{
    "Open":               reflect.ValueOf(bolt.Open),
  }
  env.PackageTypes["kv/bolt"] = map[string]reflect.Type{
    "DB":                 reflect.TypeOf(bolt.DB{}),
    "Cursor":             reflect.TypeOf(bolt.Cursor{}),
    "Bucket":             reflect.TypeOf(bolt.Bucket{}),
    "Options":            reflect.TypeOf(bolt.Options{}),
    "Tx":                 reflect.TypeOf(bolt.Tx{}),
  }
}
