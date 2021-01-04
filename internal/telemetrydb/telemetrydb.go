package telemetrydb

import (
  "database/sql"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
	_ "github.com/mattn/go-sqlite3"
)

var TDB *sql.DB

func Telemetrydb_Init() *sql.DB {
  var err error

  TDB, err = sql.Open("sqlite3", conf.TelemetryDB)
  if err != nil {
    return nil
  }
  for _, s := range TDB_CREATE {
    err = TDBExec(s)
    if err != nil {
      return nil
    }
  }
  for _, i := range TDB_IX_CREATE {
    err = TDBExec(i)
    if err != nil {
      return nil
    }
  }
  return TDB
}

func TDBExec(sql string) error {
  if TDB != nil {
    _, err := TDB.Exec(sql)
    if err != nil {
      return err
    }
  } else {
    return nil
  }
  return nil
}

func TDBVersion() string {
  if TDB != nil {
    rows, err := TDB.Query(`select sqlite_version()`)
    if err == nil {
      for rows.Next() {
        var version string
        rows.Scan(&version)
        rows.Close()
        return version
      }
    }
  }
  return "UNKNOWN"
}

func TDBCount(table int, key string) (res int, err error) {
  res = 0
  var stmt *sql.Stmt
  if TDB != nil {
    if table == 0 {
      stmt, err = TDB.Prepare("select count(*) from Counters where key = ?")
    } else if table == 1 {
      stmt, err = TDB.Prepare("select count(*) from Metric where key = ?")
    } else if table == 2 {
      stmt, err = TDB.Prepare("select count(*) from History where key = ?")
    } else {
      stmt, err = TDB.Prepare("select count(*) from History where key = ? ")
    }
    if err != nil {
      return
    } else {
      err = stmt.QueryRow(key).Scan(&res)
    }
  }
  return
}

func TDBValue(table int, key string) (res interface{}, err error) {
  res = 0
  if TDB != nil {
    var stmt *sql.Stmt
    if table == 0 {
      stmt, err = TDB.Prepare("select value from Counters where key = ? order by timestamp desc limit 1")
    } else if table == 1 {
      stmt, err = TDB.Prepare("select value from Metric where key = ? order by timestamp desc limit 1")
    } else if table == 2 {
      stmt, err = TDB.Prepare("select value from History where key = ? order by timestamp desc limit 1")
    } else {
      stmt, err = TDB.Prepare("select value from History where key = ? order by timestamp desc limit 1")
    }
    if err != nil {
      return
    } else {
      if stmt != nil {
        err = stmt.QueryRow(key).Scan(&res)
      }
    }
  }
  return
}


func TelemetrydbHousekeeping(n int) (before int64, after int64, merr error) {
  var keys []string
  var key string
  var stmt *sql.Stmt
  keys = make([]string, 0)
  if TDB != nil {
    stmt, err := TDB.Prepare("select count(*) from History")
    if err != nil {
      return
    }
    err = stmt.QueryRow().Scan(&before)
    stmt.Close()
    if err != nil {
      return
    }
    rows, err := TDB.Query(`select distinct key from History`)
    if err != nil {
      merr = err
      return
    }
    for rows.Next() {
      rows.Scan(&key)
      keys = append(keys, key)
    }
    rows.Close()
    for _, dkey := range keys {
      tx, err := TDB.Begin()
      if err != nil {
        merr = err
        return
      }
      dstmt, err := tx.Prepare(`delete from History where id <= (select id from (select id from History where key = ? order by timestamp desc limit 1 offset ?))`)
      if err != nil {
        merr = err
        return
      }
      _, err = dstmt.Exec(dkey, n)
      if err != nil {
        dstmt.Close()
        merr = err
        return
      }
      tx.Commit()
      dstmt.Close()
    }
  }
  stmt, err := TDB.Prepare("select count(*) from History")
  if err != nil {
    merr = err
    return
  }
  err = stmt.QueryRow().Scan(&after)
  stmt.Close()
  return
}

func Telemetrydb_Fin() {
  if TDB != nil {
    TDB.Close()
  }
}
