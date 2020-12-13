package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  humanize "github.com/dustin/go-humanize"
)

func init() {
  env.Packages["convert/humanize"] = map[string]reflect.Value{
    "Bytes":            reflect.ValueOf(humanize.Bytes),
    "Time":             reflect.ValueOf(humanize.Time),
    "Ordinal":          reflect.ValueOf(humanize.Ordinal),
    "Comma":            reflect.ValueOf(humanize.Comma),
    "Ftoa":             reflect.ValueOf(humanize.Ftoa),
    "SI":               reflect.ValueOf(humanize.SI),
  }
  env.PackageTypes["convert/humanize"] = map[string]reflect.Type{

  }
}
