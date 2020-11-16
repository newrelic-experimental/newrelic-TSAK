package clips

import (
  "fmt"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/si"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
)

func EvalClips(ch int, name string, fun string) bool {
  _out := gabs.New()
  _out.Set(name, "name")
  _out.Set(nr.NowMillisec(), "timestamp")
  _out.Set(si.SysInfo().Hostname, "hostname")
  _out.Set("tsak", "evtSource")
  _out.Set(true, "isClips")
  ret, err := Env().Eval(fun)
  if err != nil {
    log.Error(fmt.Sprintf("CLIPS.eval.error = %v", err))
    return false
  }
  _out.Set(ret, "value")
  piping.To(ch, []byte(_out.String()))
  return true
}

func EvalRet(fun string) interface{} {
  ret, err := Env().Eval(fun)
  if err != nil {
    log.Error(fmt.Sprintf("CLIPS.eval.error = %v", err))
    return false
  }
  return ret
}
