package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  sender "github.com/blacked/go-zabbix"
)

func init() {
  env.Packages["protocols/zabbix/sender"] = map[string]reflect.Value{
    "New":              reflect.ValueOf(sender.NewSender),
    "Packet":           reflect.ValueOf(sender.NewPacket),
    "M":                reflect.ValueOf(sender.NewMetric),
  }
  env.PackageTypes["protocols/zabbix/sender"] = map[string]reflect.Type{
    "Metric":           reflect.TypeOf(sender.Metric{}),
  }
}
