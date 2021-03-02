package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  stat "gonum.org/v1/gonum/stat"
  floats "gonum.org/v1/gonum/floats"
)

func NumNorm(data []float64) []float64 {
  res := make([]float64, len(data))
  xmin := floats.Min(data)
  xmax := floats.Max(data)
  diff := xmax-xmin
  if diff == 0 {
    for i:=0;i<len(data);i++ {
      res[i] = 0.0
    }
  } else {
    for i:=0;i<len(data);i++ {
      res[i] = (data[i] - xmin)/diff
    }
  }
  return res
}

func NumStand(data []float64) []float64 {
  xmean := stat.Mean(data, nil)
  xdev  := stat.StdDev(data, nil)
  res := make([]float64, len(data))
  for i:=0;i<len(data);i++ {
    res[i] = (data[i] - xmean)/xdev
  }
  return res
}


func init() {
  env.Packages["num/norm"] = map[string]reflect.Value{
    "Normalize":                reflect.ValueOf(NumNorm),
    "Standartize":              reflect.ValueOf(NumStand),
  }
  env.PackageTypes["num/norm"] = map[string]reflect.Type{

  }
}
