package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  stat "gonum.org/v1/gonum/stat"
)

func StatSmoothMovingAverageWeightless(p []float64, win int64) []float64 {
  var c, w, pmin, pmax int64
  var m float64
  res := make([]float64, len(p))
  c = 0
  pmin = 0
  pmax = 0
  for c <= int64(len(p)) {
    w = win
    pmin = c - w
    if pmin < 0 {
      w = 1
      pmin = c
    }
    pmax = c + w
    if pmax > int64(len(p)) {
        pmax = int64(len(p))
        w = int64(len(p)) - c
        pmin = c - w
    }
    subp := p[pmin:pmax]
    if len(subp) == 0 {
      c += 1
      continue
    }
    m = stat.Mean(subp, nil)
    res[c] = m
    c += 1
  }
  return res
}

func StatSmoothWeightless(p []float64, win int64) []float64 {
  var c, w, pmin, pmax int64
  var m float64
  res := make([]float64, len(p))
  c = 0
  pmin = 0
  pmax = 0
  for range p {
    w = win
    pmin = c - w
    if pmin < 0 {
      pmin = 0
    }
    pmax = c + w
    if pmax > int64(len(p)) {
        pmax = int64(len(p))
    }
    subp := p[pmin:pmax]
    if len(subp) == 0 {
      c += 1
      continue
    }
    m = stat.Mean(subp, nil)
    res[c] = m
    c += 1
  }
  return res
}

func init() {
  env.Packages["num/stat"] = map[string]reflect.Value{
    "Bhattacharyya":                reflect.ValueOf(stat.Bhattacharyya),
    "ChiSquare":                    reflect.ValueOf(stat.ChiSquare),
    "CircularMean":                 reflect.ValueOf(stat.CircularMean),
    "Correlation":                  reflect.ValueOf(stat.Correlation),
    "Covariance":                   reflect.ValueOf(stat.Covariance),
    "CrossEntropy":                 reflect.ValueOf(stat.CrossEntropy),
    "Entropy":                      reflect.ValueOf(stat.Entropy),
    "ExKurtosis":                   reflect.ValueOf(stat.ExKurtosis),
    "GeometricMean":                reflect.ValueOf(stat.GeometricMean),
    "HarmonicMean":                 reflect.ValueOf(stat.HarmonicMean),
    "Hellinger":                    reflect.ValueOf(stat.Hellinger),
    "Histogram":                    reflect.ValueOf(stat.Histogram),
    "JensenShannon":                reflect.ValueOf(stat.JensenShannon),
    "Kendall":                      reflect.ValueOf(stat.Kendall),
    "KolmogorovSmirnov":            reflect.ValueOf(stat.KolmogorovSmirnov),
    "KullbackLeibler":              reflect.ValueOf(stat.KullbackLeibler),
    "LinearRegression":             reflect.ValueOf(stat.LinearRegression),
    "Mean":                         reflect.ValueOf(stat.Mean),
    "MeanStdDev":                   reflect.ValueOf(stat.MeanStdDev),
    "MeanVariance":                 reflect.ValueOf(stat.MeanVariance),
    "Mode":                         reflect.ValueOf(stat.Mode),
    "Moment":                       reflect.ValueOf(stat.Moment),
    "MomentAbout":                  reflect.ValueOf(stat.MomentAbout),
    "Skew":                         reflect.ValueOf(stat.Skew),
    "SortWeighted":                 reflect.ValueOf(stat.SortWeighted),
    "SortWeightedLabeled":          reflect.ValueOf(stat.SortWeightedLabeled),
    "StdDev":                       reflect.ValueOf(stat.StdDev),
    "StdErr":                       reflect.ValueOf(stat.StdErr),
    "StdScore":                     reflect.ValueOf(stat.StdScore),
    "Variance":                     reflect.ValueOf(stat.Variance),
    "Smooth":                       reflect.ValueOf(StatSmoothMovingAverageWeightless),
    "SmoothStatic":                       reflect.ValueOf(StatSmoothWeightless),
  }
  env.PackageTypes["num/stat"] = map[string]reflect.Type{

  }
}
