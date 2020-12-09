package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  zbxagent "github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix/agent"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix/client"
)

func init() {
  env.Packages["zabbix/agent"] = map[string]reflect.Value{
    "New":              reflect.ValueOf(zbxagent.NewAgent),
  }
  env.PackageTypes["zabbix/agent"] = map[string]reflect.Type{
    "Agent":             reflect.TypeOf(zbxagent.Agent{}),
    "Packet":            reflect.TypeOf(client.Packet{}),
  }
}
