package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  zabbix "github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix"
)

func init() {
  env.Packages["protocols/zabbix/client"] = map[string]reflect.Value{
    "Get":           reflect.ValueOf(zabbix.GetAnItemFromZabbix),
  }
  env.PackageTypes["protocols/zabbix/client"] = map[string]reflect.Type{

  }
}
