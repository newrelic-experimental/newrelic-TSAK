package tsak

import (
  "fmt"
  "os"
  "syscall"
  "time"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/pidfile"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

func StartWithPid() {
  var err error
  pidfile.Pidfile = fmt.Sprintf("%s/%s.pid", conf.AppPath, conf.ID)
  log.Trace(fmt.Sprintf("PID(start): %v", pidfile.Pidfile))
  err = pidfile.Write()
  if err != nil {
    log.Error(fmt.Sprintf("PID file write error: %v", err))
  }
}

func StopWithPid() {
  var err error
  pidfile.Pidfile = fmt.Sprintf("%s/%s.pid", conf.AppPath, conf.ID)
  log.Trace(fmt.Sprintf("PID(stop): %v", pidfile.Pidfile))
  pid, err := pidfile.Read()
  if err != nil {
    log.Error(fmt.Sprintf("PID file read error: %v", err))
    return
  }
  proc, err := os.FindProcess(pid)
  if err != nil {
    log.Error(fmt.Sprintf("Can not find the process with PID=%v: %v", pid, err))
    return
  }
  if proc != nil {
    log.Trace(fmt.Sprintf("Sending signal TERM to PID=%v", pid))
    err = proc.Signal(syscall.SIGTERM)
    if err != nil {
      log.Error(fmt.Sprintf("Error sending signal TERM to PID=%v: %v", pid, err))
      return
    }
    N := 0
    for N < 10 {
      time.Sleep(time.Second)
      log.Trace(fmt.Sprintf("[%v] Searching for process PID=%v", N, pid))
      err := proc.Signal(syscall.SIGTERM)
      if err != nil {
        os.Remove(pidfile.Pidfile)
        log.Trace(fmt.Sprintf("Process been killed PID=%v", pid))
        return
      }
      N++
    }
    proc, err := os.FindProcess(pid)
    if err != nil {
      os.Remove(pidfile.Pidfile)
      log.Trace(fmt.Sprintf("Process finally gave up PID=%v", pid))
      return
    }
    log.Trace(fmt.Sprintf("Sending signal KILL to PID=%v", pid))
    err = proc.Signal(syscall.SIGKILL)
    if err != nil {
      log.Error(fmt.Sprintf("Error sending signal KILL to PID=%v: %v", pid, err))
      return
    }
    os.Remove(pidfile.Pidfile)
  }
}
