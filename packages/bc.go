package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "git.mills.io/prologic/bitcask"
)

func init() {
  env.Packages["bc"] = map[string]reflect.Value{
    "Open":                reflect.ValueOf(bitcask.Open),
  }
  env.PackageTypes["bc"] = map[string]reflect.Type{
    "Bitcask":          reflect.TypeOf(bitcask.Bitcask{}),
  }
}
