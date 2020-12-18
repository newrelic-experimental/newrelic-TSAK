package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/alitto/pond"
)



func init() {
  env.Packages["stdlib/pond"] = map[string]reflect.Value{
    "IdleTimeout":        reflect.ValueOf(pond.IdleTimeout),
    "MinWorkers":         reflect.ValueOf(pond.MinWorkers),
    "PanicHandler":       reflect.ValueOf(pond.PanicHandler),
    "Strategy":           reflect.ValueOf(pond.Strategy),
    "RatedResizer":       reflect.ValueOf(pond.RatedResizer),
    "Eager":              reflect.ValueOf(pond.Eager),
    "Balanced":           reflect.ValueOf(pond.Balanced),
    "Lazy":               reflect.ValueOf(pond.Lazy),
    "New":                reflect.ValueOf(pond.New),
  }
  env.PackageTypes["stdlib/pond"] = map[string]reflect.Type{
    "WorkerPool":         reflect.TypeOf(pond.WorkerPool{}),
    "TaskGroup":          reflect.TypeOf(pond.TaskGroup{}),

  }
}
