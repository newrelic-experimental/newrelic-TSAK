package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/buger/jsonparser"
)


func init() {
  env.Packages["json/parser"] = map[string]reflect.Value{
    "Get":                reflect.ValueOf(jsonparser.Get),
    "GetInt":             reflect.ValueOf(jsonparser.GetInt),
    "GetString":          reflect.ValueOf(jsonparser.GetString),
    "GetFloat":           reflect.ValueOf(jsonparser.GetFloat),
    "GetBoolean":         reflect.ValueOf(jsonparser.GetBoolean),
    "ArrayEach":          reflect.ValueOf(jsonparser.ArrayEach),
    "EachKey":            reflect.ValueOf(jsonparser.EachKey),
  }
  env.PackageTypes["json/parser"] = map[string]reflect.Type{

  }
}
