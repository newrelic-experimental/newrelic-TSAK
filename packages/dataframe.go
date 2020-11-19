package packages

import (
  "fmt"
  "os"
  "reflect"
  "context"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  dataframe "github.com/rocketlaunchr/dataframe-go"
  dfimport "github.com/rocketlaunchr/dataframe-go/imports"

)

func DFFromCsv(fname string) *dataframe.DataFrame {
  ctx := context.TODO()
  f, err := os.Open(fname)
  if err != nil {
    log.Trace(fmt.Sprintf("DF.open.error: %s", err))
    return nil
  }
  defer f.Close()
  df, err := dfimport.LoadFromCSV(ctx, f, dfimport.CSVLoadOptions{
    LargeDataSet: true,
    InferDataTypes: true,
  })
  if err != nil {
    log.Trace(fmt.Sprintf("DF.import.error: %s", err))
  }
  return df
}

func init() {
  env.Packages["dataframe"] = map[string]reflect.Value{
    "Series":                         reflect.ValueOf(dataframe.NewSeriesFloat64),
    "NewSeriesInt64":                 reflect.ValueOf(dataframe.NewSeriesInt64),
    "NewSeriesFloat64":               reflect.ValueOf(dataframe.NewSeriesFloat64),
    "NewSeriesString":                reflect.ValueOf(dataframe.NewSeriesString),
    "NewSeriesTime":                  reflect.ValueOf(dataframe.NewSeriesTime),
    "New":                            reflect.ValueOf(dataframe.NewDataFrame),
    "SeriesIdx":                      reflect.ValueOf(dataframe.SeriesIdx),
    "SeriesName":                     reflect.ValueOf(dataframe.SeriesName),
    "FromCSV":                        reflect.ValueOf(DFFromCsv),
  }
  env.PackageTypes["dataframe"] = map[string]reflect.Type{
    "SeriesFloat64":          reflect.TypeOf(dataframe.SeriesFloat64{}),
    "SeriesInt64":            reflect.TypeOf(dataframe.SeriesInt64{}),
    "SeriesString":           reflect.TypeOf(dataframe.SeriesString{}),
    "SeriesTime":             reflect.TypeOf(dataframe.SeriesTime{}),
    "DataFrame":              reflect.TypeOf(dataframe.DataFrame{}),
    "ValuesOptions":          reflect.TypeOf(dataframe.ValuesOptions{}),
  }
}
