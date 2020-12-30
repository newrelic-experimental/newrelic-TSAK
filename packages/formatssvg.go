package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/ajstarks/svgo"
)

func init() {
  env.Packages["formats/svg"] = map[string]reflect.Value{
    "New":          reflect.ValueOf(svg.New),
  }
  env.PackageTypes["formats/svg"] = map[string]reflect.Type{
    "SVG":          reflect.TypeOf(svg.SVG{}),
  }
}
