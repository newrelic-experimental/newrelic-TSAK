package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/antchfx/htmlquery"
  "golang.org/x/net/html"
)

func init() {
  env.Packages["formats/html"] = map[string]reflect.Value{
    "Parse":            reflect.ValueOf(htmlquery.Parse),
    "LoadURL":          reflect.ValueOf(htmlquery.LoadURL),
    "LoadDoc":          reflect.ValueOf(htmlquery.LoadDoc),
    "Find":             reflect.ValueOf(htmlquery.Find),
    "FindOne":          reflect.ValueOf(htmlquery.FindOne),
    "SelectAttr":       reflect.ValueOf(htmlquery.SelectAttr),
  }
  env.PackageTypes["formats/html"] = map[string]reflect.Type{
    "Node":             reflect.TypeOf(html.Node{}),
  }
}
