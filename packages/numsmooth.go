package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "gonum.org/v1/gonum/dsp/fourier"
)



func init() {
  env.Packages["num/ft/dst"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(fourier.NewDST),
  }
  env.PackageTypes["num/ft/dct"] = map[string]reflect.Type{
    "DST":           reflect.TypeOf(fourier.DST{}),
  }
}
