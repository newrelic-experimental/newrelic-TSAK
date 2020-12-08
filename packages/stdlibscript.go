package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/bitfield/script"
)


func init() {
  env.Packages["stdlib/script"] = map[string]reflect.Value{
    "File":                  reflect.ValueOf(script.File),
    "Stdin":                 reflect.ValueOf(script.Stdin),
    "Echo":                  reflect.ValueOf(script.Echo),
    "Exec":                  reflect.ValueOf(script.Exec),
    "IfExists":              reflect.ValueOf(script.IfExists),
    "FindFiles":             reflect.ValueOf(script.FindFiles),
    "ListFiles":             reflect.ValueOf(script.ListFiles),
    "Slice":                 reflect.ValueOf(script.Slice),
  }
  env.PackageTypes["stdlib/script"] = map[string]reflect.Type{

  }
}
