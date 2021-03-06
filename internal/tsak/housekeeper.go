package tsak

import (
  "fmt"
  "sync"
  "time"
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/script"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/cron"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
)

var HOUSE_EVERY = (1 * time.Second)
var REPORT_EVERY = 15

func HouseProc() {
  var start = nr.NowMillisec()
  cron.Start()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    housekeeper()
    log.Trace("Housekeeper thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("Housekeeper() duration", start)
  }(signal.WG())
}

func HouseShutdown() {
  log.Trace("Housekeeper terminating")
  cron.Stop()
}

func housekeeper() {
  if conf.Conf != "" {
    res := script.RunScript("house", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Housekeeper()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  c := 0
  for ! signal.ExitRequested() {
    time.Sleep(HOUSE_EVERY)
    cron.Run()
    if c > REPORT_EVERY {
      log.Trace("Running housekeeper")
      housekeeperReport()
      if conf.House != "" {
        script.RunScript("house", conf.House)
      }
      if conf.Hkeep > 0 {
        before, after, err := telemetrydb.TelemetrydbHousekeeping(conf.Hkeep)
        if err != nil {
          log.Error(fmt.Sprintf("Housekeeper returned an error: %v", err))
        } else {
          log.Trace(fmt.Sprintf("Housekeeper before: %v, after: %v", before, after))
          telemetrydb.Metric("tsak.TDB.HOUSEKEEPER.before", before)
          telemetrydb.Metric("tsak.TDB.HOUSEKEEPER.after", after)
        }
      }
      c = 0
    } else {
      c += 1
    }
  }
  signal.ExitRequest()
}

func housekeeperReport() {
  telemetrydb.Metric("tsak.INCH.size", piping.Len(piping.INCH))
  telemetrydb.Metric("tsak.OUTCH.size", piping.Len(piping.OUTCH))
  telemetrydb.Metric("tsak.CLIPS.size", piping.Len(piping.CLIPS))
  telemetrydb.Metric("tsak.FACTS.size", piping.Len(piping.FACTS))
  telemetrydb.Metric("tsak.EVAL.size", piping.Len(piping.EVAL))
  if conf.MetricsToNR {
    nr.RecordValue("tsak.INCH.size", "Number of elements in TSAK pipes", piping.Len(piping.INCH))
    nr.RecordValue("tsak.OUTCH.size", "Number of elements in TSAK pipes", piping.Len(piping.OUTCH))
  }
}
