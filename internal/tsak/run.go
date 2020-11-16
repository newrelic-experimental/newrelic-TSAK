package tsak

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
)

func Run() {
  nr.RecordEvidence("Run() checkpoint is reached")
  if conf.Run != "" {
    signal.Reserve(1)
    nr.RecordEvidence("Exclsive Run() checkpoint reached")
    RunProc()
    nr.RecordEvidence("End of exclsive Run() checkpoint reached")
  } else {
    if conf.In != "" {
      nr.RecordEvidence("Running In() code")
      signal.Reserve(1)
      InProc()
    }
    if conf.Proc != "" {
      nr.RecordEvidence("Running Proc() code")
      signal.Reserve(1)
      ProcProc()
    }
    if conf.Out != "" {
      nr.RecordEvidence("Running Proc() code")
      signal.Reserve(1)
      OutProc()
    }
    if conf.Clips != "" {
      nr.RecordEvidence("Running Clips() code as Main")
      piping.To(piping.CLIPS, []byte(fmt.Sprintf(`(batch* "%s")`, conf.Clips)))
    }
  }
  nr.RecordEvidence("Running Housekeeping() code")
  signal.Reserve(1)
  HouseProc()
  signal.Loop()
}
