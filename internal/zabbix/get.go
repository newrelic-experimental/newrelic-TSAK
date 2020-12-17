package zabbix

import (
  "time"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
  zg "github.com/vulogov/go-zabbix-get/zabbix"
)

func GetAnItemFromZabbix(host string, key string, timeout int64) (res interface{}, err error) {
  _res, err := zg.Get(host, key, time.Duration(timeout) * time.Second)
  res = stdlib.ToValue(_res)
  return
}
