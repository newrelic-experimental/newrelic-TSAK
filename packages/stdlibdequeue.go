package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/gammazero/deque"
)



func init() {
  env.Packages["stdlib/deque"] = map[string]reflect.Value{

  }
  env.PackageTypes["stdlib/deque"] = map[string]reflect.Type{
    "Deque":              reflect.TypeOf(deque.Deque{}),
  }
}
