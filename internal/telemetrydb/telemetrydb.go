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
  for _, i := range TDB_IX_CREATE {
    err = TDBExec(i)
    if err != nil {
      log.Error(fmt.Sprintf("TDB index initialization failure: %v", err))
    }
  }
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

func Telemetrydb_Fin() {
  log.Trace(fmt.Sprintf("TelemetryDB closes at: %v", conf.TelemetryDB))
  if TDB != nil {
    TDB.Close()
  }
}
