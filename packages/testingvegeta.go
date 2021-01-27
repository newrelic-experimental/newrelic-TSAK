package packages

import (
  "time"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/tsenart/vegeta/lib"
)

func VegetaAttackRateGet(url string, _rate int, _duration int) (got uint64, want uint64) {
  var hits uint64

  tr    := vegeta.NewStaticTargeter(vegeta.Target{Method: "GET", URL: url})
	rate  := vegeta.Rate{Freq: _rate, Per: time.Second}
	atk   := vegeta.NewAttacker()
  for range atk.Attack(tr, rate, time.Duration(_duration)*time.Second, "") {
		hits++
	}
  got, want = hits, uint64(rate.Freq)
  return
}


func init() {
  env.Packages["testing/vegeta"] = map[string]reflect.Value{
    "Get":              reflect.ValueOf(VegetaAttackRateGet),
  }
}
