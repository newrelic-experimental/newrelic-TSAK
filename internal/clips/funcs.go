package clips

import (
    "fmt"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

var clipsFun = make(map[string]interface{})

func AddClipsFun(name string, f interface{}) {
  clipsFun[name] = f
}

func InitFunctions() {
  log.Trace("CLIPS functions definitions reached")
  RegisterFunctions()
  for k, fun := range clipsFun {
    log.Trace(fmt.Sprintf("CLIPS.fun %s", k))
    err := env.DefineFunction(k, fun)
    if err != nil {
      log.Error(fmt.Sprintf("CLIPS.fun.error: %v", err))
    }
  }
}
