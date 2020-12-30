package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  keys "github.com/cossacklabs/themis/gothemis/keys"
)

func init() {
  env.Packages["crypto/themis/keys"] = map[string]reflect.Value{
    "TypeRSA":                reflect.ValueOf(keys.TypeRSA),
    "TypeEC":                 reflect.ValueOf(keys.TypeEC),
    "Assymmetric":            reflect.ValueOf(keys.New),
    "Symmetric":              reflect.ValueOf(keys.NewSymmetricKey),
    "New":                    reflect.ValueOf(keys.New),
  }
  env.PackageTypes["crypto/themis/keys"] = map[string]reflect.Type{
    "SymmetricKey":           reflect.TypeOf(keys.SymmetricKey{}),
    "PublicKey":              reflect.TypeOf(keys.PublicKey{}),
    "PrivateKey":             reflect.TypeOf(keys.PrivateKey{}),
    "Keypair":                reflect.TypeOf(keys.Keypair{}),
  }
}
