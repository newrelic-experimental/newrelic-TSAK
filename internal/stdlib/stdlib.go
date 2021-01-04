package stdlib

import (
  "strconv"
  "time"
)

func NowMilliseconds() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func ToValue(repr string) interface{} {
  vb, err := strconv.ParseBool(repr)
  if err == nil {
    return vb
  }
  vi, err := strconv.ParseInt(repr, 10, 64)
  if err == nil {
    return vi
  }
  vf, err := strconv.ParseFloat(repr, 64)
  if err == nil {
    return vf
  }
  return repr
}
