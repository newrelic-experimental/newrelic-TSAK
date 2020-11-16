package tsak

import (
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
)

func TSAK() {
  Init()
  Run()
  signal.Loop()
  Fin()
}
