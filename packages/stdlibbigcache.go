package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/allegro/bigcache/v2"
)


func init() {
  env.Packages["stdlib/bigcache"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(bigcache.NewBigCache),
    "DefaultConfig":      reflect.ValueOf(bigcache.DefaultConfig),
    "Expired":            reflect.ValueOf(bigcache.Expired),
    "NoSpace":            reflect.ValueOf(bigcache.NoSpace),
    "Deleted":            reflect.ValueOf(bigcache.Deleted),
  }
  env.PackageTypes["stdlib/bigcache"] = map[string]reflect.Type{
    "BigCache":           reflect.TypeOf(bigcache.BigCache{}),
    "Config":             reflect.TypeOf(bigcache.Config{}),
    "Response":           reflect.TypeOf(bigcache.Response{}),
    "Stats":              reflect.TypeOf(bigcache.Stats{}),
  }
}
