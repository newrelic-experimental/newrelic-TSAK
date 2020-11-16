package packages

import (
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "reflect"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["tlog"] = map[string]reflect.Value{
    "Trace":    reflect.ValueOf(log.Trace),
    "Info":     reflect.ValueOf(log.Info),
    "Error":    reflect.ValueOf(log.Error),
    "Warning":  reflect.ValueOf(log.Warning),
    "Event":    reflect.ValueOf(log.Event),
    "SendEvent":reflect.ValueOf(nr.SendEvent),
    "Metric":   reflect.ValueOf(log.Metric),
  }
}
