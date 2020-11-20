package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/ringbuffer"
)




func init() {
  env.Packages["stdlib/ringbuffer"] = map[string]reflect.Value{
    "New":             reflect.ValueOf(ringbuffer.New),
  }
  env.PackageTypes["stdlib/ringbuffer"] = map[string]reflect.Type{
    "Ring":           reflect.TypeOf(ringbuffer.Ring{}),
  }
}
