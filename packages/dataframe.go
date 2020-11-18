package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  dataframe "github.com/rocketlaunchr/dataframe-go"
)

func init() {
  env.Packages["dataframe"] = map[string]reflect.Value{
    "Series":                         reflect.ValueOf(dataframe.NewSeriesFloat64),
    "NewSeriesInt64":                 reflect.ValueOf(dataframe.NewSeriesInt64),
    "NewSeriesFloat64":               reflect.ValueOf(dataframe.NewSeriesFloat64),
    "NewSeriesString":                reflect.ValueOf(dataframe.NewSeriesString),
    "New":                            reflect.ValueOf(dataframe.NewDataFrame),
    "SeriesIdx":                      reflect.ValueOf(dataframe.SeriesIdx),
    "SeriesName":                     reflect.ValueOf(dataframe.SeriesName),
  }
  env.PackageTypes["dataframe"] = map[string]reflect.Type{
    "SeriesFloat64":          reflect.TypeOf(dataframe.SeriesFloat64{}),
    "SeriesInt64":            reflect.TypeOf(dataframe.SeriesInt64{}),
    "SeriesString":           reflect.TypeOf(dataframe.SeriesString{}),
    "DataFrame":              reflect.TypeOf(dataframe.DataFrame{}),
    "ValuesOptions":          reflect.TypeOf(dataframe.ValuesOptions{}),
  }
}
