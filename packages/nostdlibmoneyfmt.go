package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/leekchan/accounting"
)

func init() {
  env.Packages["nostdlib/finance/fmt"] = map[string]reflect.Value{
    "Default":             reflect.ValueOf(accounting.DefaultAccounting),
    "New":                 reflect.ValueOf(accounting.NewAccounting),
  }
  env.PackageTypes["nostdlib/finance/fmt"] = map[string]reflect.Type{
    "Accounting":          reflect.TypeOf(accounting.Accounting{}),
    "Locale":              reflect.TypeOf(accounting.Locale{}),
  }
}
