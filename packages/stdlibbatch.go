package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/batch"
)


func init() {
  env.Packages["stdlib/batch"] = map[string]reflect.Value{
    "Configuration":         reflect.ValueOf(batch.Configuration),
  }
  env.PackageTypes["stdlib/batch"] = map[string]reflect.Type{

  }
}
