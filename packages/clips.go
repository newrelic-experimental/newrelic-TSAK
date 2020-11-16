package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/clips"
)

func init() {
  env.Packages["clips"] = map[string]reflect.Value{
    "Eval":               reflect.ValueOf(clips.EvalClips),
    "eval":               reflect.ValueOf(clips.EvalRet),
  }
}
