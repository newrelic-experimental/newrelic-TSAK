package piping

import (
  "time"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  zmq "github.com/pebbe/zmq4"
)

func RRBroker(fe, be string) {
	frontend := Router(fe)
  backend := Dealer(be)
  if frontend == nil && backend == nil {
    log.Error("REQUEST/REPLY broker sockets are not ready")
    return
  }
	frontend.Bind(fe)
	backend.Bind(be)

	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

  log.Trace("Entered ReqRepBroker")
	for ! signal.ExitRequested() {
		sockets, _ := poller.Poll(time.Duration(conf.Timeout)*time.Second)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				for ! signal.ExitRequested() {
					msg, err := s.Recv(zmq.DONTWAIT)
          if err != nil {
            continue
          }
					if more, _ := s.GetRcvmore(); more {
						backend.Send(msg, zmq.SNDMORE)
					} else {
						backend.Send(msg, 0)
						break
					}
				}
			case backend:
				for ! signal.ExitRequested() {
					msg, err := s.Recv(zmq.DONTWAIT)
          if err != nil {
            continue
          }
					if more, _ := s.GetRcvmore(); more {
						frontend.Send(msg, zmq.SNDMORE)
					} else {
						frontend.Send(msg, 0)
						break
					}
				}
			}
		}
	}
  log.Trace("Exit ReqRepBroker")
}
