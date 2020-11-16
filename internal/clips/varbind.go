package clips

import (
    "fmt"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
    "github.com/keysight/clipsgo/pkg/clips"
)

var vars = make(map[clips.Symbol]interface{})

func SetVarBind(name clips.Symbol, v interface{}) bool {
  log.Trace(fmt.Sprintf("CLIPS.varbind[%s] = %v", name, v))
  vars[name] = v
  return true
}

func GetVarBind(name clips.Symbol) interface{} {
  if val, ok := vars[name]; ok {
    return val
  }
  return false
}

func GetVarBindDef(name clips.Symbol, _default interface{}) interface{} {
  if val, ok := vars[name]; ok {
    return val
  }
  SetVarBind(name, _default)
  return _default
}
