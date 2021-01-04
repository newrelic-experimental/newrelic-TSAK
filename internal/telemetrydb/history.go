package telemetrydb

import (
  "fmt"
  "errors"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)

func TDBHistoryAdd(key string, pkt *gabs.Container) bool {
  var stamp = stdlib.NowMilliseconds()
  var data = pkt.String()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      log.Error(fmt.Sprintf("Error building history/add transaction:", err))
      return false
    }
    stmt, err := tx.Prepare("insert into History(timestamp,key,value) values(?,?,?)")
    if err != nil {
      log.Error(fmt.Sprintf("Error building history/add query:", err))
    } else {
      _, err = stmt.Exec(stamp, key, data)
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

func TDBHistoryLast(key string) (res *gabs.Container, err error) {
  var pkt []byte
  n, err := TDBCount(2, key)
  if n == 0 {
    return nil, errors.New(fmt.Sprintf("No history %v", key))
  }
  dbpkt, err := TDBValue(2, key)
  pkt = dbpkt.([]byte)
  if err != nil {
    log.Error(fmt.Sprintf("Error executing histroy/last query:", err))
    return
  }
  res, err = gabs.ParseJSON(pkt)
  if err != nil {
    log.Error(fmt.Sprintf("Error parsing history/last result:", err))
  }
  return
}

func TDBHistoryGet(key string) (res *gabs.Container, err error) {
  res = gabs.New()
  res.Set(conf.TelemetryDB, "source")
  res.Array("data")
  stmt, err := TDB.Prepare("select timestamp,value from History where key = ? order by timestamp desc")
  if err != nil {
    log.Error(fmt.Sprintf("Error building history/get query:", err))
    return
  }
  rows, err := stmt.Query(key)
  if err != nil {
    log.Error(fmt.Sprintf("Error executing history/get query:", err))
    return
  }
  for rows.Next() {
    var stamp int64
    var pkt string
    rows.Scan(&stamp, &pkt)
    dpkt, err := gabs.ParseJSON([]byte(pkt))
    if err != nil {
      continue
    }
    dpkt.Set(key, "key")
    dpkt.Set(stamp, "timestamp")
    res.ArrayAppend(dpkt.Data(), "data")
  }
  rows.Close()
  stmt.Close()
  return
}
