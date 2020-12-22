package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/tcpserver"
)

func init() {
  env.Packages["protocols/tcp/server"] = map[string]reflect.Value{
    "New":           reflect.ValueOf(tcpserver.New),
    "TLSNew":           reflect.ValueOf(tcpserver.NewWithTLS),
  }
  env.PackageTypes["protocols/tcp/server"] = map[string]reflect.Type{
    "Client":           reflect.TypeOf(tcpserver.Client{}),
    "Server":           reflect.TypeOf(tcpserver.AServer{}),
  }
}
