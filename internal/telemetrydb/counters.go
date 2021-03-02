package telemetrydb

import (
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
)

func TDBCounterAdd(key string) (bool,error) {
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      return false, err
    }
    stmt, err := tx.Prepare("insert into Counters(timestamp,key,value) values(?,?,1) on conflict(key) do update set timestamp=?, value=value+1")
    if err != nil {
      return false, err
    } else {
      _, err = stmt.Exec(stdlib.NowMilliseconds(), key, stdlib.NowMilliseconds())
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

func Counter(key string) (bool,error) {
  return TDBCounterAdd(key)
}

func TDBCounterGet(key string) (res int64, err error) {
  dbres, err := TDBValue(0, key)
  res = dbres.(int64)
  return
}
