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

func ProcProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    procproc()
    log.Trace("Proc thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("Proc() duration", start)
  }(signal.WG())
}

func procproc() {
  if conf.Conf != "" {
    res := script.RunScript("proc", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Proc()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  script.RunScript("proc", conf.Proc)
}
