package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/Jeffail/gabs"
)

func JSONParsing(data []byte) *gabs.Container {
  c, err := gabs.ParseJSON(data)
  if err != nil {
    return gabs.New()
  }
  return c
}

func JSONPipeline(n int64) chan *gabs.Container {
  return make(chan *gabs.Container, n)
}

func JSONPipelineMap(n int64) chan map[string]interface{} {
  return make(chan map[string]interface{}, n)
}

func init() {
  env.Packages["djson"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(gabs.New),
    "Parse":              reflect.ValueOf(JSONParsing),
    "Pipe":               reflect.ValueOf(JSONPipeline),
    "MapPipe":            reflect.ValueOf(JSONPipelineMap),
  }
  env.PackageTypes["djson"] = map[string]reflect.Type{
    "Container":          reflect.TypeOf(gabs.Container{}),
  }
}
