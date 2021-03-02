package packages

import (
  "bytes"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/wcharczuk/go-chart"
)

func ChartSeries(n int) []interface{} {
  return make([]interface{}, n)
}

func ChartBasic(title string, series []chart.Series) (out []byte, err error) {
  graph := chart.Chart{
    Title: title,
    Series: series,
  }

  buffer := bytes.NewBuffer([]byte{})
  err = graph.Render(chart.PNG, buffer)
  out = buffer.Bytes()
  return
}

func init() {
  env.Packages["utils/charts"] = map[string]reflect.Value{
    "Chart":              reflect.ValueOf(ChartBasic),
    "Series":             reflect.ValueOf(ChartSeries),
  }
  env.PackageTypes["utils/charts"] = map[string]reflect.Type{
    "X":                  reflect.TypeOf(chart.XAxis{}),
    "Y":                  reflect.TypeOf(chart.YAxis{}),
    "ContinuousRange":    reflect.TypeOf(chart.ContinuousRange{}),
    "ContinuousSeries":   reflect.TypeOf(chart.ContinuousSeries{}),
  }
}
