package packages

import (
  "github.com/elastic/go-sysinfo/types"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/si"
)


func init() {
  env.Packages["sysinfo"] = map[string]reflect.Value{
    "Host":               reflect.ValueOf(si.SysInfo),
  }
  env.PackageTypes["sysinfo"] = map[string]reflect.Type{
    "HostInfo":           reflect.TypeOf(types.HostInfo{}),
  }
}
