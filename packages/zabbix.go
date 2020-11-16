package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  zbxproxy "github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix/proxy"
)

func init() {
  env.Packages["zabbix"] = map[string]reflect.Value{
    "Proxy":      reflect.ValueOf(zbxproxy.NewProxy),
  }
  env.PackageTypes["zabbix"] = map[string]reflect.Type{
    "Proxy":                reflect.TypeOf(zbxproxy.Proxy{}),
    "ProxyConfig":          reflect.TypeOf(zbxproxy.ProxyConfig{}),
    "ProxyResponse":        reflect.TypeOf(zbxproxy.ProxyResponse{}),
    "ProxyConfigResponse":  reflect.TypeOf(zbxproxy.ProxyConfigResponse{}),
    "ProxyConfigDiscovered":  reflect.TypeOf(zbxproxy.ProxyConfigDiscovered{}),
  }
}
