package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/goccy/go-graphviz"
)

func init() {
  env.Packages["formats/gv"] = map[string]reflect.Value{
    "New":              reflect.ValueOf(graphviz.New),
    "PNG":              reflect.ValueOf(graphviz.PNG),
    "JPG":              reflect.ValueOf(graphviz.JPG),
    "SVG":              reflect.ValueOf(graphviz.SVG),
  }
  env.PackageTypes["formats/gv"] = map[string]reflect.Type{
    "Graphviz":          reflect.TypeOf(graphviz.Graphviz{}),
  }
}
