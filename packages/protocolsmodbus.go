package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/goburrow/modbus"
)


func init() {
  env.Packages["protocols/modbus"] = map[string]reflect.Value{
    "TCP":               reflect.ValueOf(modbus.TCPClient),
    "Serial":            reflect.ValueOf(modbus.RTUClient),
  }
  env.PackageTypes["protocols/modbus"] = map[string]reflect.Type{

  }
}
