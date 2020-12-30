package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/modem"
)


func init() {
  env.Packages["protocols/modem"] = map[string]reflect.Value{
    "New":               reflect.ValueOf(modem.New),
  }
  env.PackageTypes["protocols/modem"] = map[string]reflect.Type{
    "Modem":             reflect.TypeOf(modem.GSMModem{}),
  }
}
