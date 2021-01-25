package snmpdb

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  nrisnmp "github.com/newrelic-experimental/nri-snmpdb/nrisnmp"
)

func Init() {
  nrisnmp.Init()
  log.Trace(fmt.Sprintf("Default location of the nrisnmp.db: %v", nrisnmp.DBName()))
}
