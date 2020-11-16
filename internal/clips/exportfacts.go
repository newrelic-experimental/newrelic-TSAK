package clips

import (
    "fmt"
    "github.com/keysight/clipsgo/pkg/clips"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
    "github.com/Jeffail/gabs"
)

func procImpliedFact(f clips.Fact) string {
  var res []interface{}
  _out := gabs.New()
  f.Extract(&res)
  c := 0
  _out.Set(f.Template().Name(), "name")
  for _, _v := range res {
    _out.Set(_v, fmt.Sprintf("value%d", c))
    c += 1
  }
  return _out.String()
}

func procTemplatedFact(f clips.Fact) string {
  var res map[string]interface{}
  _out := gabs.New()
  f.Extract(&res)
  _out.Set(f.Template().Name(), "name")
  _out.Set(f.Asserted(), "asserted")
  for k, v := range res {
    _out.Set(v, k)
  }
  return _out.String()
}

func ExportAllFacts(ch int) bool {
  var res string
  log.Trace(fmt.Sprintf("Exporting all facts to %d", ch))
  for _, f := range env.Facts() {
    if f.Template().Implied() {
      res = procImpliedFact(f)
      piping.To(ch, []byte(res))
    } else {
      res = procTemplatedFact(f)
      piping.To(ch, []byte(res))
    }
  }
  return true
}

func ExportAssertedFacts(ch int) bool {
  var res string
  log.Trace(fmt.Sprintf("Exporting all facts to %d", ch))
  for _, f := range env.Facts() {
    if f.Asserted() {
      if f.Template().Implied() {
        res = procImpliedFact(f)
        piping.To(ch, []byte(res))
      } else {
        res = procTemplatedFact(f)
        piping.To(ch, []byte(res))
      }
    }
  }
  return true
}
