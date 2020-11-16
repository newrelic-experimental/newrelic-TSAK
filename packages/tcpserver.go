package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/orkunkaraduman/go-tcpserver"
)


func init() {
  env.Packages["net/tcpserver"] = map[string]reflect.Value{

  }
  env.PackageTypes["net/tcpserver"] = map[string]reflect.Type{
    "Server":          reflect.TypeOf(tcpserver.TCPServer{}),
  }
}
