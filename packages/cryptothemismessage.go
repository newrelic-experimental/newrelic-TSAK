package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  message "github.com/cossacklabs/themis/gothemis/message"
)

func init() {
  env.Packages["crypto/themis/message"] = map[string]reflect.Value{
    "New":                    reflect.ValueOf(message.New),
  }
  env.PackageTypes["crypto/themis/keys"] = map[string]reflect.Type{
    "SecureMessage":          reflect.TypeOf(message.SecureMessage{}),
  }
}
