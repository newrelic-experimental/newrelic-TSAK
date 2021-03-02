package script

import (
  "fmt"
  "io/ioutil"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/mattn/anko/env"
  _ "github.com/mattn/anko/packages"
  _ "github.com/newrelic-experimental/newrelic-TSAK/packages"
	"github.com/mattn/anko/vm"
)

var e = make(map[string]*env.Env)

func Define(name string, key string, value interface{}) {
  err := e[name].Define(key, value)
  if err != nil {
    log.Error(fmt.Sprintf("Def(%[1]s) = %[2]s", key, err))
  }
}


func InitScript() {
  log.Trace("Initialize internal script engine")
  e["in"] = env.NewEnv()
  e["proc"] = env.NewEnv()
  e["out"] = env.NewEnv()
  e["house"] = env.NewEnv()

  for k := range env.Packages {
    log.Trace(fmt.Sprintf("Module: %[1]s", k))
  }

  for k, _ := range e {
    Define(k, "ANSWER", 42)
    Define(k, "ID", conf.ID)
    Define(k, "VMNAME", k)
  }

}

func RunString(ename string, script string) string {
  if _, ok := e[ename]; ok {
    res, err := vm.Execute(e[ename], nil, script)
    if err != nil {
      log.Error(fmt.Sprintf("Error executing %[1]s", script))
      fmt.Println(err)
      return "FAIL"
    }
    return fmt.Sprintf("%v", res)
  } else {
    log.Error(fmt.Sprintf("No VM registered %[1]s", ename))
    return "FAIL"
  }
}

func RunScript(ename string, fname string) string {
  if fname == "" {
    log.Error(fmt.Sprintf("Script file not specified"))
    return "FAIL"
  }
  log.Trace(fmt.Sprintf("Running %[1]s", fname))
  buf, err := ioutil.ReadFile(fname)
  if err != nil {
    log.Error(fmt.Sprintf("Error reading %[1]s: %[2]s", fname, err))
    return "FAIL"
  }
  script := string(buf)

  if _, ok := e[ename]; ok {
    res, err := vm.Execute(e[ename], nil, script)
    if err != nil {
      log.Error(fmt.Sprintf("Error executing %[1]s", fname))
      fmt.Println(err)
      return "FAIL"
    }
    return fmt.Sprintf("%v", res)
  } else {
    log.Error(fmt.Sprintf("No VM registered %[1]s", ename))
    return "FAIL"
  }
}
