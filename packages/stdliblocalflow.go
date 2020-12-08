package packages

import (
  "fmt"
  "context"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  flow "github.com/kamildrazkiewicz/go-flow"
  hunch "github.com/aaronjan/hunch"
)

func StdlibFlowAll(execs ...hunch.Executable) ([]interface{}, error) {
  ctx := context.Background()
  r, err := hunch.All(ctx, execs...)
  if err != nil {
    log.Trace(fmt.Sprintf("Flow(All) error: %s", err))
  }
  return r, err
}

func StdlibFlowFirst(first int, execs ...hunch.Executable) ([]interface{}, error) {
  ctx := context.Background()
  r, err := hunch.Take(ctx, first, execs...)
  if err != nil {
    log.Trace(fmt.Sprintf("Flow(First) error: %s", err))
  }
  return r, err
}

func StdlibFlowLast(last int, execs ...hunch.Executable) ([]interface{}, error) {
  ctx := context.Background()
  r, err := hunch.Last(ctx, last, execs...)
  if err != nil {
    log.Trace(fmt.Sprintf("Flow(Last) error: %s", err))
  }
  return r, err
}

func StdlibFlowPipe(execs ...hunch.ExecutableInSequence) (interface{}, error) {
  ctx := context.Background()
  r, err := hunch.Waterfall(ctx, execs...)
  if err != nil {
    log.Trace(fmt.Sprintf("Flow(Pipe) error: %s", err))
  }
  return r, err
}

func StdlibFlowRetry(count int, exec hunch.Executable) (interface{}, error) {
  ctx := context.Background()
  r, err := hunch.Retry(ctx, count, exec)
  if err != nil {
    log.Trace(fmt.Sprintf("Flow(Retry) error: %s", err))
  }
  return r, err
}

func init() {
  env.Packages["stdlib/flow"] = map[string]reflect.Value{
    "New":                  reflect.ValueOf(flow.New),
    "All":                  reflect.ValueOf(StdlibFlowAll),
    "First":                reflect.ValueOf(StdlibFlowFirst),
    "Last":                 reflect.ValueOf(StdlibFlowLast),
    "Pipe":                 reflect.ValueOf(StdlibFlowPipe),
    "Retry":                reflect.ValueOf(StdlibFlowRetry),
  }
  env.PackageTypes["stdlib/flow"] = map[string]reflect.Type{
    "Flow":                   reflect.TypeOf(flow.Flow{}),
  }
}
