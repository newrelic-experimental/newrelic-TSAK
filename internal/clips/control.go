package clips

import (
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

var DoFact = true
var DoCmd = true

func EnableCmdPipe() bool {
  DoCmd = true
  log.Trace("ENABLING CMD PIPE")
  return true
}

func DisableCmdPipe() bool {
  DoCmd = false
  log.Trace("DISABLING CMD PIPE")
  return false
}

func EnableFactPipe() bool {
  DoFact = true
  log.Trace("ENABLING FACTS PIPE")
  return true
}

func DisableFactPipe() bool {
  DoFact = false
  log.Trace("DISABLING FACTS PIPE")
  return false
}
