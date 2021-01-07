package tsak

import (
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/clips"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
)

func Fin() {
  P2PShutdown()
  telemetrydb.Telemetrydb_Fin()
  piping.Shutdown()
  clips.Shutdown()
  log.Shutdown()
  HouseShutdown()
  log.Event("TsakEvent", logrus.Fields{
    "message":    "Application exited",
    "evtc":       1,
  })
}
