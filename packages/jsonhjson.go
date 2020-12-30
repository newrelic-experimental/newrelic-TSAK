package packages

import (
  "reflect"
  "encoding/json"
  "github.com/Jeffail/gabs"
  "github.com/mattn/anko/env"
  "github.com/hjson/hjson-go"
)

func HjsonParse(data []byte) (res *gabs.Container, err error) {
  var hdata map[string]interface{}

  res = nil
  hjson.Unmarshal(data, &hdata)
  b, err := json.Marshal(hdata)
  if err != nil {
    return
  }
  res, err = gabs.ParseJSON(b)
  return
}

func init() {
  env.Packages["json/hjson"] = map[string]reflect.Value{
    "Parse":                reflect.ValueOf(HjsonParse),
  }
  env.PackageTypes["json/hjson"] = map[string]reflect.Type{

  }
}
