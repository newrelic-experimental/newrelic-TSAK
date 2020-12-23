package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/piquette/finance-go"
  "github.com/piquette/finance-go/quote"
)

func init() {
  env.Packages["nostdlib/finance/quote"] = map[string]reflect.Value{
    "Ticker":             reflect.ValueOf(quote.Get),
  }
  env.PackageTypes["nostdlib/finance/quote"] = map[string]reflect.Type{
    "Quote":              reflect.TypeOf(finance.Quote{}),
  }
}
