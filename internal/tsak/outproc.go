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

func OutProc() {
  var start = nr.NowMillisec()
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    outproc()
    log.Trace("Out thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("Out() duration", start)
  }(signal.WG())
}

func outproc() {
  if conf.Conf != "" {
    res := script.RunScript("out", conf.Conf)
    log.Event(
      "Bootstrap is loaded for Out()",
      logrus.Fields{
        "result":     res,
        "confSource": conf.Conf,
    })
  }
  script.RunScript("out", conf.Out)
}
