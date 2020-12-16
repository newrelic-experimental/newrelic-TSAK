package packages

import (
  "fmt"
  "time"
  "reflect"
  "strconv"
  "context"
  "github.com/google/uuid"
  "github.com/mattn/anko/env"
  "github.com/erikdubbelboer/gspt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/cron"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/teris-io/shortid"
  "golang.org/x/sync/semaphore"
)


func NowMilliseconds() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func UUID() string {
  uid, _ := uuid.NewUUID()
  return uid.String()
}

func String(src []byte) string {
  return string(src)
}
func StdlibBytes(src string) []byte {
  return []byte(src)
}


func Int(n string) int32 {
  res, err := strconv.Atoi(n)
  if err != nil {
    return 0
  } else {
    return int32(res)
  }
}

func StdlibToValue(repr string) interface{} {
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

func StdlibShortID() *shortid.Shortid {
  sid, err := shortid.New(1, shortid.DefaultABC, uint64(time.Now().UTC().UnixNano()))
  if err != nil {
    log.Error(fmt.Sprintf("UniqID.error: %s", err))
    return nil
  }
  return sid
}

func StdlibSemaphore(n uint64) *semaphore.Weighted {
  return semaphore.NewWeighted(int64(n))
}

func StdlibTODO() context.Context {
  return context.TODO()
}

func init() {
  env.Packages["stdlib"] = map[string]reflect.Value{
    "Answer":         reflect.ValueOf(42),
    "SetProcTitle":   reflect.ValueOf(gspt.SetProcTitle),
    "ExitRequest":    reflect.ValueOf(signal.ExitRequest),
    "ExitRequested":  reflect.ValueOf(signal.ExitRequested),
    "Release":        reflect.ValueOf(signal.Release),
    "NowMilliseconds":reflect.ValueOf(NowMilliseconds),
    "Cron":           reflect.ValueOf(cron.AddToCron),
    "UUID":           reflect.ValueOf(UUID),
    "From":           reflect.ValueOf(piping.From),
    "To":             reflect.ValueOf(piping.To),
    "Len":            reflect.ValueOf(piping.Len),
    "INCH":           reflect.ValueOf(piping.INCH),
    "OUTCH":          reflect.ValueOf(piping.OUTCH),
    "CLIPS":          reflect.ValueOf(piping.CLIPS),
    "FACTS":          reflect.ValueOf(piping.FACTS),
    "EVAL":           reflect.ValueOf(piping.EVAL),
    "String":         reflect.ValueOf(String),
    "Bytes":          reflect.ValueOf(StdlibBytes),
    "Query":          reflect.ValueOf(nr.Query),
    "Int":            reflect.ValueOf(Int),
    "UniqID":         reflect.ValueOf(StdlibShortID),
    "Semaphore":      reflect.ValueOf(StdlibSemaphore),
    "TODO":           reflect.ValueOf(StdlibTODO),
    "ToValue":        reflect.ValueOf(StdlibToValue),
  }
  env.PackageTypes["stdlib"] = map[string]reflect.Type{
    "WeightedSemaphore":     reflect.TypeOf(semaphore.Weighted{}),
  }
}
