package batch

import (
  "encoding/json"
  "github.com/Jeffail/gabs"
  "github.com/hjson/hjson-go"
)

func Configuration(data string) (res *gabs.Container, err error) {
  var hdata map[string]interface{}

  res = nil
  hjson.Unmarshal([]byte(data), &hdata)
  b, err := json.Marshal(hdata)
  if err != nil {
    return
  }
  res, err = gabs.ParseJSON(b)
  return
}

func ParseJobs(cfg *gabs.Container) () {
  
}
