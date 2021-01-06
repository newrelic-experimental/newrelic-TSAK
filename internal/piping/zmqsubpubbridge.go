package piping

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  zmq "github.com/pebbe/zmq4"
)

func PUBSUBBridge(fe, be string) {
	frontend := XSubscriber(fe)
  backend := XPublisher(be)
  if frontend == nil && backend == nil {
    log.Error("SUB/PUB broker sockets are not ready")
    return
  }
	frontend.Connect(fe)
	backend.Bind(be)

	pxyctl := ZS("pub:proxycontrol")
  log.Trace(fmt.Sprintf("Entering PUB/SUB proxy %v->%v", fe, be))
  zmq.ProxySteerable(frontend, backend, nil, pxyctl)
}
