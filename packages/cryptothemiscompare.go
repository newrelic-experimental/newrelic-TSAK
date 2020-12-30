package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  compare "github.com/cossacklabs/themis/gothemis/compare"
)

func init() {
  env.Packages["crypto/themis/compare"] = map[string]reflect.Value{
    "New":                    reflect.ValueOf(compare.New),
    "Match":                  reflect.ValueOf(compare.Match),
    "NoMatch":                reflect.ValueOf(compare.NoMatch),
    "NotReady":               reflect.ValueOf(compare.NotReady),
  }
  env.PackageTypes["crypto/themis/compare"] = map[string]reflect.Type{
    "SecureCompare":          reflect.TypeOf(compare.SecureCompare{}),
  }
}
