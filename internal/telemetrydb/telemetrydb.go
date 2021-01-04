package telemetrydb

import (
  "database/sql"
	"fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
	_ "github.com/mattn/go-sqlite3"
)

var TDB *sql.DB

func Telemetrydb_Init() *sql.DB {
  var err error
  log.Trace(fmt.Sprintf("TelemetryDB opens at: %v", conf.TelemetryDB))
  TDB, err = sql.Open("sqlite3", conf.TelemetryDB)
  if err != nil {
    log.Error(fmt.Sprintf("TelemetryDB opens failure: %v", err))
    return nil
  }
  log.Trace(fmt.Sprintf("SQLITE version is %v", TDBVersion()))
  for _, s := range TDB_CREATE {
    err = TDBExec(s)
    if err != nil {
      log.Error(fmt.Sprintf("TDB initialization failure: %v", err))
    }
  }
  log.Trace("TDB tables created")
  for _, i := range TDB_IX_CREATE {
    err = TDBExec(i)
    if err != nil {
      log.Error(fmt.Sprintf("TDB index initialization failure: %v", err))
    }
  }
  log.Trace("TDB indexes created")
  return TDB
}

func TDBExec(sql string) error {
  if TDB != nil {
    _, err := TDB.Exec(sql)
    if err != nil {
      log.Error(fmt.Sprintf("TelemetryDB exec failure on \"%v\": %v", sql, err))
      return err
    }
  } else {
    log.Error("Access to an unititialized TelemetryDB")
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
      log.Error(fmt.Sprintf("Error building count query:", err))
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
      log.Error(fmt.Sprintf("Error building value query:", err))
    } else {
      if stmt != nil {
        err = stmt.QueryRow(key).Scan(&res)
      } else {
        log.Error("Attempt to call an empty statemement")
      }
    }
  }
  return
}


func TelemetrydbHousekeeping(n int) (before int64, after int64) {
  var keys []string
  var key string
  keys = make([]string, 0)
  if TDB != nil {
    var stmt *sql.Stmt
    stmt, err := TDB.Prepare("select count(*) from History")
    if err != nil {
      log.Error(fmt.Sprintf("Error building #1 housekeeping query: %v", err))
      return
    }
    err = stmt.QueryRow().Scan(&before)
    stmt.Close()
    if err != nil {
      log.Error(fmt.Sprintf("Error discovering housekeeping state: %v", err))
      return
    }
    rows, err := TDB.Query(`select distinct key from History`)
    if err != nil {
      log.Error(fmt.Sprintf("Error discovering keys in housekeeper: %v", err))
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
        log.Error(fmt.Sprintf("Error initiating housekeeper transaction: %v", err))
        return
      }
      dstmt, err := tx.Prepare(`delete from History where id <= (select id from (select id from History where key = ? order by timestamp desc limit 1 offset ?))`)
      if err != nil {
        log.Error(fmt.Sprintf("Error preparing housekeeper query: %v", err))
        return
      }
      _, err = dstmt.Exec(dkey, n)
      if err != nil {
        log.Error(fmt.Sprintf("Error executing housekeeper transaction: %v", err))
      }
      tx.Commit()
      dstmt.Close()
    }
  }
  stmt, err := TDB.Prepare("select count(*) from History")
  if err != nil {
    log.Error(fmt.Sprintf("Error building #2 housekeeping query: %v", err))
    return
  }
  err = stmt.QueryRow().Scan(&after)
  stmt.Close()
  log.Trace(fmt.Sprintf("Housekeeper finished. Before %v, after %v", before, after))
  return
}

func Telemetrydb_Fin() {
  log.Trace(fmt.Sprintf("TelemetryDB closes at: %v", conf.TelemetryDB))
  if TDB != nil {
    TDB.Close()
  }
}
