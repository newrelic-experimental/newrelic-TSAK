package telemetrydb

import (
  "fmt"
  "errors"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
)

func Metric(key string, value interface{}) (bool, error) {
  var pkt = gabs.New()
  pkt.Set(value, "value")
  return TDBMetricAdd(key, pkt)
}

func TDBMetricAdd(key string, pkt *gabs.Container) (bool, error) {
  var stamp = stdlib.NowMilliseconds()
  var data = pkt.String()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      return false, err
    }
    stmt, err := tx.Prepare("insert into Metric(timestamp,key,value) values(?,?,?) on conflict(key) do update set timestamp=?, value=?")
    if err != nil {
      return false, err
    } else {
      _, err = stmt.Exec(stamp, key, data, stamp, data)
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

func TDBMetricGet(key string) (res *gabs.Container, err error) {
  var pkt []byte
  n, err := TDBCount(1, key)
  if n == 0 {
    return nil, errors.New(fmt.Sprintf("No metric %v", key))
  }
  dbpkt, err := TDBValue(1, key)
  pkt = dbpkt.([]byte)
  if err != nil {
    return
  }
  res, err = gabs.ParseJSON(pkt)
  return
}
