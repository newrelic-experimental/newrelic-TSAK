package signal

import (
  "os"
  "fmt"
  "os/signal"
  "syscall"
  "sync"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 128)
var ng = 0

func signalHandler() {
  log.Trace("Running signal handler")
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
  <- c
  log.Info(fmt.Sprintf("Interruption signal detected. N=%d", ng))
  for i := 0; i < ng; i++ {
    exitChan <- true
  }
}

func ExitRequest() {
  exitChan <- true
}

func ExitRequested() bool {
  if len(exitChan) > 0 {
    log.Trace(fmt.Sprintf("Exit requested: %[1]d",len(exitChan)))
  }
  select {
  case _, ok := <- exitChan:
    if ok {
      ExitRequest()
      return true
    } else {
      ExitRequest()
      return true
    }
  default:
    return false
  }
}

func WG() *sync.WaitGroup {
  return &wg
}

func Len() int {
  return len(exitChan)
}

func InitSignal() {
  log.Trace("Configuring signals")
  go signalHandler()
}

func Reserve(n int) {
  ng = ng + n
  wg.Add(n)
}

func Release(n int) {
  ng = ng - n
  for i := 0; i < n; i++ {
    wg.Done()
  }
}

func Loop() {
  wg.Wait()
  // for ! ExitRequested() {
  //   stdlib.SleepForASecond()
  // }
}
