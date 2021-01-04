package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
)

func init() {
  env.Packages["telemetrydb"] = map[string]reflect.Value{
    "Version":                reflect.ValueOf(telemetrydb.TDBVersion),
    "Counter":                reflect.ValueOf(telemetrydb.TDBCounterGet),
    "Increase":               reflect.ValueOf(telemetrydb.TDBCounterAdd),
    "Add":                    reflect.ValueOf(telemetrydb.TDBMetricAdd),
    "Get":                    reflect.ValueOf(telemetrydb.TDBMetricGet),
  }
  env.PackageTypes["telemetrydb"] = map[string]reflect.Type{
  }
}
