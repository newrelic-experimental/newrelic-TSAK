package packages

import (
  "reflect"
  "gopkg.in/h2non/gentleman.v2"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["protocols/http/client"] = map[string]reflect.Value{
    "New":                    reflect.ValueOf(gentleman.New),  
  }
  env.PackageTypes["protocols/http/client"] = map[string]reflect.Type{
    "Client":                 reflect.TypeOf(gentleman.Client{}),
    "Request":                reflect.TypeOf(gentleman.Request{}),
    "Response":               reflect.TypeOf(gentleman.Response{}),
  }
}
