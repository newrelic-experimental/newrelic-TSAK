package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  zmq "github.com/pebbe/zmq4"
)



func init() {
  env.Packages["stdlib/zmq"] = map[string]reflect.Value{
    "Publisher":          reflect.ValueOf(piping.Publisher),
    "Subscriber":         reflect.ValueOf(piping.Subscriber),
    "XPublisher":         reflect.ValueOf(piping.XPublisher),
    "XSubscriber":        reflect.ValueOf(piping.XSubscriber),
    "Req":                reflect.ValueOf(piping.Req),
    "Rep":                reflect.ValueOf(piping.Rep),
    "RepReqBroker":       reflect.ValueOf(piping.RRBroker),
    "MaxSockets":         reflect.ValueOf(zmq.GetMaxSockets),
    "DONTWAIT":           reflect.ValueOf(zmq.DONTWAIT),

  }
  env.PackageTypes["stdlib/zmq"] = map[string]reflect.Type{
    "Socket":         reflect.TypeOf(zmq.Socket{}),
  }
}
