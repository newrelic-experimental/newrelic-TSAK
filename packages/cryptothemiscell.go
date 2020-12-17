package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  cell "github.com/cossacklabs/themis/gothemis/cell"
)

func init() {
  env.Packages["crypto/themis/cell"] = map[string]reflect.Value{
    "WithKey":                reflect.ValueOf(cell.SealWithKey),
    "WithPassword":           reflect.ValueOf(cell.SealWithPassphrase),
  }
  env.PackageTypes["crypto/themis/cell"] = map[string]reflect.Type{
    "SecureCell":             reflect.TypeOf(cell.SecureCell{}),
  }
}
