package tsak

import (
  "sync"
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/script"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)

func RunProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    runproc()
    log.Trace("Run thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("Run() duration", start)
  }(signal.WG())
}

func runproc() {
  if conf.Conf != "" {
    res := script.RunScript("proc", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Run()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  script.RunScript("proc", conf.Run)
}
