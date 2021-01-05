package piping

import (
  "fmt"
  "bytes"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
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

func ClearZMQINTRR () {
  log.Trace("Disabling RetryAfterEINTR for ZMQ core")
  zmq.SetRetryAfterEINTR(false)
  zmqCtx.SetRetryAfterEINTR(false)
}

func To(dst int, _data []byte) {
  var data = bytes.NewBuffer(_data)

  if dst == INCH {
    telemetrydb.Counter("tsak.INCH.sent")
    pipeIn <- data.String()
  } else if dst == OUTCH {
    telemetrydb.Counter("tsak.OUTCH.sent")
    pipeOut <- data.String()
  } else if dst == CLIPS {
    telemetrydb.Counter("tsak.CLIPS.sent")
    clipsIn <- data.String()
  } else if dst == FACTS {
    telemetrydb.Counter("tsak.FACTS.sent")
    factsIn <- data.String()
  } else if dst == EVAL {
    telemetrydb.Counter("tsak.EVAL.sent")
    evalIn <- data.String()
  } else {
    telemetrydb.Counter("tsak.TO.errors")
    log.Error("Trying to send data to non-existent pipeline")
  }
}

func From(src int) []byte {
  if src == INCH && len(pipeIn) > 0 {
    telemetrydb.Counter("tsak.INCH.recv")
    return []byte(<-pipeIn)
  } else if src == OUTCH && len(pipeOut) > 0 {
    telemetrydb.Counter("tsak.OUTCH.recv")
    return []byte(<-pipeOut)
  } else if src == CLIPS && len (clipsIn) > 0 {
    telemetrydb.Counter("tsak.CLIPS.recv")
    return []byte(<-clipsIn)
  } else if src == FACTS && len (factsIn) > 0 {
    telemetrydb.Counter("tsak.FACTS.recv")
    return []byte(<-factsIn)
  } else if src == EVAL && len (evalIn) > 0 {
    telemetrydb.Counter("tsak.EVAL.recv")
    return []byte(<-evalIn)
  } else {
    telemetrydb.Counter("tsak.FROM.errors")
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

func Init() {
  log.Trace("Initializing Pipelines")
  if conf.IPv6 {
    log.Trace("IPv6 was enabled for ZMQ")
  }
  zmq.SetIpv6(conf.IPv6)
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
