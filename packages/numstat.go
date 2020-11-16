package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  stat "gonum.org/v1/gonum/stat"
)

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
  }
  env.PackageTypes["num/stat"] = map[string]reflect.Type{

  }
}
