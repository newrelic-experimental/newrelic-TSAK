package cron

import (
  rcron "github.com/robfig/cron"
)

var tcron = rcron.New()

func Start() {
  tcron.Start()
}

func Stop() {
  tcron.Stop()
}

func AddToCron(c_spec string, fun func()) {
  tcron.AddFunc(c_spec, fun)
}


func Run() {
  tcron.Run()
}
