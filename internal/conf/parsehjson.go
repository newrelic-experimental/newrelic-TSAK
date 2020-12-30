package conf

import (
  "encoding/json"
  "github.com/Jeffail/gabs"
  "github.com/hjson/hjson-go"
)

func ParseHjson(data []byte) (res *gabs.Container, err error) {
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
