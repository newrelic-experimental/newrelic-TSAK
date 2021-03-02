package telemetrydb

import (
  "fmt"
  "errors"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)

func TDBHistoryAdd(key string, pkt *gabs.Container) (bool, error) {
  var stamp = stdlib.NowMilliseconds()
  var data = pkt.String()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      return false, err
    }
    stmt, err := tx.Prepare("insert into History(timestamp,key,value) values(?,?,?)")
    if err != nil {
      return false, err
    } else {
      _, err = stmt.Exec(stamp, key, data)
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

func TDBHistoryInsert(pkt *gabs.Container) (bool, error) {
  var stamp = stdlib.NowMilliseconds()
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      return false, err
    }
    stmt, err := tx.Prepare("insert into History(timestamp,key,value) values(?,?,?)")
    if err != nil {
      return false, err
    }
    childs, err := pkt.S("data").Children()
    if err != nil {
      tx.Commit()
      stmt.Close()
      return false, err
    }
    for _, item := range childs {
      key := item.Path("key").Data()
      value := gabs.New()
      value.Set(item.Path("value").Data(), "value")
      cmap, err := item.S().ChildrenMap()
      if err != nil {
        return false, err
      }
      for attrkey, attrval := range cmap {
        if attrkey == "key" {
          continue
        }
        value.Set(attrval.Data(), attrkey)
      }
      _, err = stmt.Exec(stamp, key, value.String())
      if err != nil {
        return false, err
      }
    }
    tx.Commit()
    stmt.Close()
    return true, nil
  }
  return false, nil
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
    return
  }
  res, err = gabs.ParseJSON(pkt)
  return
}

func TDBHistoryGet(key string) (res *gabs.Container, err error) {
  res = gabs.New()
  res.Set(conf.TelemetryDB, "source")
  res.Array("data")
  stmt, err := TDB.Prepare("select timestamp,value from History where key = ? order by timestamp desc")
  if err != nil {
    return
  }
  rows, err := stmt.Query(key)
  if err != nil {
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
