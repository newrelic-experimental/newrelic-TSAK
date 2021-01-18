package packages

import (
  "time"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/tj/go-naturaldate"
)

func TimeHumanAsNow(rel string) (time.Time, error) {
  return naturaldate.Parse(rel, time.Now())
}

func init() {
  env.Packages["time/human"] = map[string]reflect.Value{
    "AsOfNow":             reflect.ValueOf(TimeHumanAsNow),
    "AsOf":                reflect.ValueOf(naturaldate.Parse),
  }
  env.PackageTypes["time/human"] = map[string]reflect.Type{
  }
}
