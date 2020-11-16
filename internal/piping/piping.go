package piping

import (
  "fmt"
  "bytes"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  // "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
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

func To(dst int, _data []byte) {
  var data = bytes.NewBuffer(_data)
  log.Trace(fmt.Sprintf("Sending %d bytes to  pipeline %d", data.Len(), dst))
  if dst == INCH {
    pipeIn <- data.String()
    log.Trace(fmt.Sprintf("%d element in pipeline IN", len(pipeIn)))
  } else if dst == OUTCH {
    pipeOut <- data.String()
    log.Trace(fmt.Sprintf("%d element in pipeline OUT", len(pipeOut)))
  } else if dst == CLIPS {
    clipsIn <- data.String()
    log.Trace(fmt.Sprintf("%d element in pipeline CLIPS", len(clipsIn)))
  } else if dst == FACTS {
    factsIn <- data.String()
    log.Trace(fmt.Sprintf("%d element in pipeline FACTS", len(factsIn)))
  } else if dst == EVAL {
    evalIn <- data.String()
    log.Trace(fmt.Sprintf("%d element in pipeline EVAL", len(factsIn)))
  } else {
    log.Error("Trying to send data to non-existent pipeline")
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
  for k, v := range zmqS {
    log.Trace(fmt.Sprintf("Closing ZMQ: %s", k))
    v.Close()
  }
  zmqCtx.Term()
  log.Trace("Pipelines are terminated")
}
