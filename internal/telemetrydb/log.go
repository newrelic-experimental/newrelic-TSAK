package telemetrydb

import (
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)

func Log(key string, message string) (bool, error) {
  if ! conf.TelemetryLog {
    return true, nil
  }
  var stamp = stdlib.NowMilliseconds()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      return false, err
    }
    stmt, err := tx.Prepare("insert into Log(timestamp,key,value) values(?,?,?)")
    if err != nil {
      return false, err
    } else {
      _, err = stmt.Exec(stamp, key, message)
      if err != nil {
        stmt.Close()
        return false, err
      } else {
        tx.Commit()
        stmt.Close()
      }
      return true, nil
    }
  }
  return false, nil
}




func TDBLogGet(key string) (res *gabs.Container, err error) {
  res = gabs.New()
  res.Set(conf.TelemetryDB, "source")
  res.Array("data")
  stmt, err := TDB.Prepare("select timestamp,value from Log where key = ? order by timestamp desc")
  if err != nil {
    return
  }
  rows, err := stmt.Query(key)
  if err != nil {
    return
  }
  for rows.Next() {
    var stamp int64
    var msg string
    rows.Scan(&stamp, &msg)
    dpkt := gabs.New()
    dpkt.Set(key, "key")
    dpkt.Set(stamp, "timestamp")
    dpkt.Set(msg, "msg")
    res.ArrayAppend(dpkt.Data(), "data")
  }
  rows.Close()
  stmt.Close()
  return
}
