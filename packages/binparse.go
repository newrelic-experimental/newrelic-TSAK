package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  bp "github.com/newrelic-experimental/newrelic-TSAK/internal/binary_pack"
)

func BinparseNew() *bp.BinaryPack {
  return new(bp.BinaryPack)
}


func init() {
  env.Packages["parse/bin"] = map[string]reflect.Value{
    "New":             reflect.ValueOf(BinparseNew),
  }
  env.PackageTypes["parse/bin"] = map[string]reflect.Type{
    "BinaryPack":          reflect.TypeOf(bp.BinaryPack{}),
  }
}
