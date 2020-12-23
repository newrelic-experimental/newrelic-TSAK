package piping

import (
  "fmt"
  "bytes"
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  humanize "github.com/dustin/go-humanize"
  zmq "github.com/pebbe/zmq4"
)

var INCH = 0
var OUTCH = 1
var CLIPS = 2
var FACTS = 3
var EVAL = 4


var pipeIn = make(chan string, 1000000)
var pipeOut = make(chan string, 1000000)
var clipsIn = make(chan string, 1000000)
var factsIn = make(chan string, 1000000)
var evalIn = make(chan string, 1000000)
var zmqS = make(map[string]*zmq.Socket)

var zmqCtx,_ = zmq.NewContext()
var zmqErr int64

var N   = make(map[int]int)
var Nb  = make(map[int]int)
var Ns  = make(map[int]int)

func To(dst int, _data []byte) {
  var data = bytes.NewBuffer(_data)
  _, ok := Nb[dst]
  if ! ok {
    Nb[dst] = 0
  }
  Nb[dst] += data.Len()
  _, ok = Ns[dst]
  if ! ok {
    Ns[dst] = 0
  }
  _, ok = N[dst]
  if ! ok {
    N[dst] = 0
  }
  if dst == INCH {
    pipeIn <- data.String()
    Ns[dst] = len(pipeIn)
  } else if dst == OUTCH {
    pipeOut <- data.String()
    Ns[dst] = len(pipeOut)
  } else if dst == CLIPS {
    clipsIn <- data.String()
    Ns[dst] = len(clipsIn)
  } else if dst == FACTS {
    factsIn <- data.String()
    Ns[dst] = len(factsIn)
  } else if dst == EVAL {
    evalIn <- data.String()
    Ns[dst] = len(evalIn)
  } else {
    log.Error("Trying to send data to non-existent pipeline")
  }
  if N[dst] > conf.Every {
    sNb := humanize.Bytes(uint64(Nb[dst]))
    log.Trace("PIPELINE statistics", logrus.Fields{
      "pipeline":       dst,
      "submitted":      N[dst],
      "bytes":          Nb[dst],
      "bytesH":         sNb,
      "elements":       Ns[dst],
    })
    N[dst]  = 0
    Ns[dst] = 0
    Nb[dst] = 0
  } else {
    N[dst] += 1
  }
}

func From(src int) []byte {
  if src == INCH && len(pipeIn) > 0 {
    return []byte(<-pipeIn)
  } else if src == OUTCH && len(pipeOut) > 0 {
    return []byte(<-pipeOut)
  } else if src == CLIPS && len (clipsIn) > 0 {
    return []byte(<-clipsIn)
  } else if src == FACTS && len (factsIn) > 0 {
    return []byte(<-factsIn)
  } else if src == EVAL && len (evalIn) > 0 {
    return []byte(<-evalIn)
  } else {
    return []byte("")
  }
}

func Len(src int) int {
  if src == INCH {
    return len(pipeIn)
  } else if src == OUTCH {
    return len(pipeOut)
  } else if src == CLIPS {
    return len(clipsIn)
  } else if src == FACTS {
    return len(factsIn)
  } else if src == EVAL {
    return len(evalIn)
  } else {
    return 0
  }
}

func Shutdown() {
  log.Trace("Terminating Pipelines")
  if zmqS != nil {
    log.Trace("Closing ZMQ sockets")
    for k, v := range zmqS {
      log.Trace(fmt.Sprintf("Closing ZMQ: %s", k))
      if v != nil {
        v.Close()
      }
    }
  }
  if zmqCtx != nil {
    log.Trace("Terminating ZMQ context")
    zmqCtx.Term()
  }
  log.Trace("Pipelines are terminated")
}
