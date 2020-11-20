package packages

import (
  "os"
  "io"
  "fmt"
  "time"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/follower"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

func FileTailNew(fname string) *follower.Follower {
  for {
    if _, err := os.Stat(fname); os.IsNotExist(err) {
      log.Trace(fmt.Sprintf("File %s does not exsis for tail", fname))
      time.Sleep(5*time.Second)
      continue
    }
    break
  }
  t, err := follower.New(fname, follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
		Reopen: true,
	})
  if err != nil {
    log.Trace(fmt.Sprintf("TAIL.error: %v", err))
    return nil
  }
  return t
}

func FileTailWatch(t *follower.Follower) []follower.Line {
  res := make([]follower.Line, 0)
  for line := range t.Lines() {
    res = append(res, line)
		if len(t.Lines()) == 0 {
      break
    }
	}
  return res
}

func init() {
  env.Packages["file/tail"] = map[string]reflect.Value{
    "New":      reflect.ValueOf(FileTailNew),
    "Watch":    reflect.ValueOf(FileTailWatch),
  }
  env.PackageTypes["file/tail"] = map[string]reflect.Type{
    "Line":           reflect.TypeOf(follower.Line{}),
    "Follower":       reflect.TypeOf(follower.Follower{}),
  }
}
