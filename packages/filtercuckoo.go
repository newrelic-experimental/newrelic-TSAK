package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/seiflotfy/cuckoofilter"
)

func init() {
  env.Packages["filter/cuckoo"] = map[string]reflect.Value{
    "New":             reflect.ValueOf(cuckoo.NewFilter),
  }
  env.PackageTypes["filter/cuckoo"] = map[string]reflect.Type{
    "Filter":          reflect.TypeOf(cuckoo.Filter{}),
  }
}
