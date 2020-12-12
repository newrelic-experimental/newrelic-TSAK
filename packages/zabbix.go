package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  zabbix "github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix"
)

func init() {
  env.Packages["protocols/zabbix"] = map[string]reflect.Value{
    "Packet":           reflect.ValueOf(zabbix.MakePacket),
    "Request":          reflect.ValueOf(zabbix.MakeReq),
    "Response":         reflect.ValueOf(zabbix.MakeResp),
    "Data":             reflect.ValueOf(zabbix.MakeData),
    "PayloadSize":      reflect.ValueOf(zabbix.GetPayloadSize),
    "ParsePacket":      reflect.ValueOf(zabbix.ParsePacket),
    "ParseRaw":         reflect.ValueOf(zabbix.ParseRaw),
    "Parse":            reflect.ValueOf(zabbix.Parse),
    "Key":              reflect.ValueOf(zabbix.ParseKey),
    "OneWay":           reflect.ValueOf(zabbix.OneWay),
    "TwoWay":           reflect.ValueOf(zabbix.TwoWay),
    "ThreeWay":         reflect.ValueOf(zabbix.ThreeWay),
  }
  env.PackageTypes["protocols/zabbix"] = map[string]reflect.Type{

  }
}
