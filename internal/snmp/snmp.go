package snmp

import (
  "fmt"
  "strings"
  "strconv"
  "github.com/hallidave/mibtool/smi"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

var mibs *smi.MIB

func InitMib(path string) {
  log.Trace(fmt.Sprintf("MIB's will be loaded from %s", path))
  mibs = smi.NewMIB(path)
}

func LoadModule(modname string) bool {
  if mibs == nil {
    log.Error("You have to initialize MIB")
    return false
  }
  err := mibs.LoadModules(modname)
  if err != nil {
    log.Error(fmt.Sprintf("Error loading %s: %s", modname, err))
    return false
  }
  log.Trace(fmt.Sprintf("Module loaded %s", modname))
  return true
}

func OID(sym string) string {
  oid, err := mibs.OID(sym)
  if err != nil {
    log.Error(fmt.Sprintf("Error resolving OID %s: %s", sym, err))
    return ""
  }
  return oid.String()
}

func IsOID(sym string) string {
  oid, err := mibs.OID(sym)
  if err != nil {
    return sym
  }
  return oid.String()
}

func SYMBOL(oid string) string {
  tmp := strings.Split(oid, ".")
  values := make([]int, 0, len(tmp))
  for _, raw := range tmp {
    v, err := strconv.Atoi(raw)
    if err != nil {
        // log.Error(fmt.Sprintf("Error converting OID element %s: %s", oid, err))
        continue
    }
    values = append(values, v)
  }
  return mibs.SymbolString(values)
}
