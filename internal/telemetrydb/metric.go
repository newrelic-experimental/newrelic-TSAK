package telemetrydb

import (
  "fmt"
  "errors"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
)

func TDBMetricAdd(key string, pkt *gabs.Container) bool {
  var stamp = stdlib.NowMilliseconds()
  var data = pkt.String()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      log.Error(fmt.Sprintf("Error building metric/add transaction:", err))
      return false
    }
    stmt, err := tx.Prepare("insert into Metric(timestamp,key,value) values(?,?,?) on conflict(key) do update set timestamp=?, value=?")
    if err != nil {
      log.Error(fmt.Sprintf("Error building metric/add query:", err))
    } else {
      _, err = stmt.Exec(stamp, key, data, stamp, data)
      if err != nil {
        stmt.Close()
        log.Error(fmt.Sprintf("Error executing metric/add query:", err))
        return false
      } else {
        tx.Commit()
        stmt.Close()
      }
      return true
    }
  }
  return false
}

func TDBMetricGet(key string) (res *gabs.Container, err error) {
  var pkt []byte
  n, err := TDBCount(1, key)
  if n == 0 {
    return nil, errors.New(fmt.Sprintf("No metric %v", key))
  }
  dbpkt, err := TDBValue(1, key)
  pkt = dbpkt.([]byte)
  if err != nil {
    log.Error(fmt.Sprintf("Error executing metric/get query:", err))
    return
  }
  res, err = gabs.ParseJSON(pkt)
  if err != nil {
    log.Error(fmt.Sprintf("Error parsing metric/get result:", err))
  }
  return
}
