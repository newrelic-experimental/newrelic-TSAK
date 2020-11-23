package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/VictoriaMetrics/fastcache"
)

var ts_cache = map[string]*fastcache.Cache{}

func StdlibFCacheNew(name string, size int) *fastcache.Cache {
  if c, ok := ts_cache[name]; ok {
    return c
  } else {
    return fastcache.New(size)
  }
}


func init() {
  env.Packages["stdlib/cache"] = map[string]reflect.Value{
    "Get":                reflect.ValueOf(StdlibFCacheNew),
  }
  env.PackageTypes["stdlib/cache"] = map[string]reflect.Type{
    "Cache":              reflect.TypeOf(fastcache.Cache{}),
  }
}
