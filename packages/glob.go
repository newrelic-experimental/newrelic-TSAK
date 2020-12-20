package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  glob "github.com/ganbarodigital/go_glob"
)

func init() {
  env.Packages["parse/globbing"] = map[string]reflect.Value{
    "New":            reflect.ValueOf(glob.NewGlob),
  }
  env.PackageTypes["parse/globbing"] = map[string]reflect.Type{
    "Glob":           reflect.TypeOf(glob.Glob{}),
  }
}
