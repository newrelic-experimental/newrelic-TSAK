package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "gonum.org/v1/gonum/dsp/fourier"
)



func init() {
  env.Packages["num/ft/dct"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(fourier.NewDCT),
  }
  env.PackageTypes["num/ft/dct"] = map[string]reflect.Type{
    "DCT":           reflect.TypeOf(fourier.DCT{}),
  }
}
