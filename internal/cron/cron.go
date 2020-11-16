package cron

import (
  "github.com/robfig/cron"
)

var tcron = cron.New()

func Start() {
  tcron.Start()
}

func Stop() {
  tcron.Stop()
}

func AddToCron(c_spec string, fun func()) {
  tcron.AddFunc(c_spec, fun)
}
