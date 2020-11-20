package packages

import (
  "fmt"
  "time"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/araddon/dateparse"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

func StdlibParsetimeSetTZ(tz string) {
  log.Trace(fmt.Sprintf("Setting local timezone as %s", tz))
  loc, err := time.LoadLocation(tz)
  if err != nil {
    log.Trace(fmt.Sprintf("Failed to set a local timezone: %s", err))
    return
  }
  time.Local = loc
}


func init() {
  env.Packages["stdlib/dateparse"] = map[string]reflect.Value{
    "TZ":             reflect.ValueOf(StdlibParsetimeSetTZ),
    "Local":          reflect.ValueOf(dateparse.ParseLocal),
    "Date":           reflect.ValueOf(dateparse.ParseAny),
    "Strict":         reflect.ValueOf(dateparse.ParseStrict),
  }
}
