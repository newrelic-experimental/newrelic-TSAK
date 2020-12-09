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
    "PayloadSize":      reflect.ValueOf(zabbix.GetPayloadSize),
    "ParsePacket":      reflect.ValueOf(zabbix.ParsePacket),
  }
  env.PackageTypes["protocols/zabbix"] = map[string]reflect.Type{

  }
}
