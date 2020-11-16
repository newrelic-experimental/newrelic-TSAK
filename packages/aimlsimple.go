package packages

import (
  // "fmt"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/goml/gobrain"
)

type SimpleFeedForward struct {
  f             gobrain.FeedForward
  isConfigured  bool
  isTrained     bool
  epoch         int
  isReport      bool
  lrate         float64
  momentum      float64
}

func SFFNew(input, hidden, output int) *SimpleFeedForward {
  sff := SimpleFeedForward{}
  sff.isConfigured = false
  sff.isTrained = false
  sff.f = gobrain.FeedForward{}
  sff.f.Init(input, hidden, output)
  return &sff
}

func (ff *SimpleFeedForward) Configure(epoch int, lrate float64, momentum float64, isReport bool) {
  ff.epoch = epoch
  ff.lrate = lrate
  ff.momentum = momentum
  ff.isReport = isReport
  ff.isConfigured = true
}

func (ff *SimpleFeedForward) Configured() bool {
  return ff.isConfigured
}

func (ff *SimpleFeedForward) Train(dataset [][][]float64) []float64 {
  ff.isTrained = true
  return ff.f.Train(dataset, ff.epoch, ff.lrate, ff.momentum, ff.isReport)
}

func (ff *SimpleFeedForward) Test(dataset [][][]float64) {
  if ff.isConfigured || ff.isTrained {
    ff.f.Test(dataset)
  }
}

func (ff *SimpleFeedForward) Update(dataset []float64) []float64 {
  if ff.isConfigured || ff.isTrained {
    return ff.f.Update(dataset)
  }
  return make([]float64, 0)
}

func init() {
  env.Packages["ai/ml/simple"] = map[string]reflect.Value{
    "New":                   reflect.ValueOf(SFFNew),
  }
  env.PackageTypes["ai/ml/simple"] = map[string]reflect.Type{
    "TrainigSet":            reflect.TypeOf([][][]float64{}),
    "FeedForward":           reflect.TypeOf(SimpleFeedForward{}),
  }
}
