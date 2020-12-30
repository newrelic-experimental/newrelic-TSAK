package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/codingsince1985/checksum"
)

func init() {
  env.Packages["crypto/hash"] = map[string]reflect.Value{
    "MD5":                reflect.ValueOf(checksum.MD5sum),
    "SHA256":             reflect.ValueOf(checksum.SHA256sum),
  }
  env.PackageTypes["crypto/hash"] = map[string]reflect.Type{

  }
}
