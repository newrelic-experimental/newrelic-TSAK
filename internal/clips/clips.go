package clips

import (
    "time"
    "sync"
    "fmt"
    "github.com/keysight/clipsgo/pkg/clips"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
)

var env = clips.CreateEnvironment()

func ReinitCLIPS() {
  env.Clear()
  env.Reset()
}

func Env() *clips.Environment {
  return env
}

func InitClips() {
  log.Trace("CLIPS module is Initializing")
  ReinitCLIPS()
  InitFunctions()
  signal.Reserve(1)
  go func(wg *sync.WaitGroup) {
    var start = nr.NowMillisec()
    defer wg.Done()
    clipsproc()
    log.Trace("CLIPS thread exiting")
    nr.RecordDuration("CLIPS() duration", start)
  }(signal.WG())
}

func Shutdown() {
  log.Trace("CLIPS module is shutting down")
  env.Delete()
}

func clipsproc() {
  for ! signal.ExitRequested() && piping.Len(piping.CLIPS) == 0 {
    time.Sleep(1*time.Second)
    if DoFact {
      for piping.Len(piping.FACTS) > 0 {
        sfact := string(piping.From(piping.FACTS))
        fact, err := env.AssertString(sfact)
        if err != nil {
          log.Error(fmt.Sprintf("CLIPS.fact.error: %v", err))
        } else {
          log.Trace(fmt.Sprintf("CLIPS.fact: %v", fact.String()))
        }
      }
    }
    if DoCmd {
      for piping.Len(piping.CLIPS) > 0 {
        cmd := string(piping.From(piping.CLIPS))
        err := env.SendCommand(cmd)
        log.Trace(fmt.Sprintf("CLIPS: %v", cmd))
        if err != nil {
          log.Error(fmt.Sprintf("CLIPS.error: %v", err))
        }
        if cmd == "(clear)" {
          log.Trace("Restoring TSAK CLIPS environment")
          InitFunctions()
        }
        if conf.Clips != "" {
          log.Trace("Main CLIPS script exit.")
          signal.ExitRequest()
        }
      }
    }
  }
}
