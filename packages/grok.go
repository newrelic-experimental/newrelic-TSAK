package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/trivago/grok"
)

func GrokNew(p map[string]string) (*grok.Grok, error) {
  g, err := grok.New(
    grok.Config{
      NamedCapturesOnly: true,
      Patterns: p,
  })
  return g, err
}


func init() {
  env.Packages["parse/grok"] = map[string]reflect.Value{
    "New":             reflect.ValueOf(GrokNew),
  }
  env.PackageTypes["parse/grok"] = map[string]reflect.Type{
    "Grok":          reflect.TypeOf(grok.Grok{}),
  }
}
