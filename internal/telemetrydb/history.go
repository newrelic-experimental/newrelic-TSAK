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

func TDBHistoryInsert(pkt *gabs.Container) bool {
  var stamp = stdlib.NowMilliseconds()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      log.Error(fmt.Sprintf("Error building history/insert transaction:", err))
      return false
    }
    stmt, err := tx.Prepare("insert into History(timestamp,key,value) values(?,?,?)")
    if err != nil {
      log.Error(fmt.Sprintf("Error building history/insert query:", err))
      return false
    }
    childs, err := pkt.S("data").Children()
    if err != nil {
      tx.Commit()
      stmt.Close()
      return false
    }
    for _, item := range childs {
      key := item.Path("key").Data()
      value := gabs.New()
      value.Set(item.Path("value").Data(), "value")
      cmap, err := item.S().ChildrenMap()
      if err != nil {
        log.Error(fmt.Sprintf("Error scan history/insert dataset:", err))
        continue
      }
      for attrkey, attrval := range cmap {
        if attrkey == "key" {
          continue
        }
        value.Set(attrval.Data(), attrkey)
      }
      _, err = stmt.Exec(stamp, key, value.String())
      if err != nil {
        log.Error(fmt.Sprintf("Error execute history/insert query:", err))
      }
    }
    tx.Commit()
    stmt.Close()
    return true
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
