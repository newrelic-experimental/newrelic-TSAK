package piping

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  zmq "github.com/pebbe/zmq4"
)

func Publisher(name string) *zmq.Socket {
  sname := fmt.Sprintf("pub:%v",name)
  s, err := zmq.NewSocket(zmq.PUB)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating publisher socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func Subscriber(name string) *zmq.Socket {
  sname := fmt.Sprintf("sub:%v",name)
  s, err := zmq.NewSocket(zmq.SUB)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating subscriber socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func XPublisher(name string) *zmq.Socket {
  sname := fmt.Sprintf("xpub:%v",name)
  s, err := zmq.NewSocket(zmq.XPUB)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating xpublisher socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func XSubscriber(name string) *zmq.Socket {
  sname := fmt.Sprintf("xsub:%v",name)
  s, err := zmq.NewSocket(zmq.XSUB)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating xsubscriber socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func Req(name string) *zmq.Socket {
  sname := fmt.Sprintf("req:%v",name)
  s, err := zmq.NewSocket(zmq.REQ)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating req socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func Rep(name string) *zmq.Socket {
  sname := fmt.Sprintf("rep:%v",name)
  s, err := zmq.NewSocket(zmq.REP)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating rep socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func Router(name string) *zmq.Socket {
  sname := fmt.Sprintf("router:%v",name)
  s, err := zmq.NewSocket(zmq.ROUTER)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating router socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}

func Dealer(name string) *zmq.Socket {
  sname := fmt.Sprintf("dealer:%v",name)
  s, err := zmq.NewSocket(zmq.DEALER)
  if err != nil {
    log.Error(fmt.Sprintf("Error creating dealer socket: %v", err))
    return nil
  }
  Z(sname, s)
  return s
}
