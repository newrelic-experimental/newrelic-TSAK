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
    "Set":                    reflect.ValueOf(telemetrydb.TDBMetricAdd),
    "Get":                    reflect.ValueOf(telemetrydb.TDBMetricGet),
    "Add":                    reflect.ValueOf(telemetrydb.TDBHistoryAdd),
    "Last":                   reflect.ValueOf(telemetrydb.TDBHistoryLast),
    "History":                reflect.ValueOf(telemetrydb.TDBHistoryGet),
    "Insert":                 reflect.ValueOf(telemetrydb.TDBHistoryInsert),
    "Housekeeper":            reflect.ValueOf(telemetrydb.TelemetrydbHousekeeping),
  }
  env.PackageTypes["telemetrydb"] = map[string]reflect.Type{
  }
}
