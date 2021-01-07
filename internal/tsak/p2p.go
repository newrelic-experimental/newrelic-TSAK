package tsak

import (
  "fmt"
  "time"
  "sync"
  "net"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
  "github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

var P2Pnode *noise.Node

func InitP2P() {
  var err error
  if conf.IsP2P {
    log.Trace("Initializing P2P layer")
    for _, b := range conf.P2PBootstrap {
      log.Trace(fmt.Sprintf("Bootstrap from node: %v", b))
    }
    if conf.P2PExternalAddress != "" {
      P2Pnode, err = noise.NewNode(
        noise.WithNodeBindHost(net.ParseIP(conf.P2PBind)),
        noise.WithNodeBindPort(uint16(conf.P2PBindPort)),
        noise.WithNodeAddress(conf.P2PExternalAddress),
        noise.WithNodeIdleTimeout(time.Duration(conf.P2PIdleTimeout)*time.Second),
        noise.WithNodeMaxInboundConnections(uint(conf.P2PMaxInbound)),
        noise.WithNodeMaxOutboundConnections(uint(conf.P2PMaxOutbound)),
      )
    } else {
      P2Pnode, err = noise.NewNode(
        noise.WithNodeBindHost(net.ParseIP(conf.P2PBind)),
        noise.WithNodeBindPort(uint16(conf.P2PBindPort)),
        noise.WithNodeIdleTimeout(time.Duration(conf.P2PIdleTimeout)*time.Second),
        noise.WithNodeMaxInboundConnections(uint(conf.P2PMaxInbound)),
        noise.WithNodeMaxOutboundConnections(uint(conf.P2PMaxOutbound)),
      )
    }
    if err == nil {
      log.Trace(fmt.Sprintf("P2P Node created: %v", P2Pnode.ID().Address))
    } else {
      log.Error(fmt.Sprintf("P2P Node creation failure: %v", err))
    }
  } else {
    log.Trace("P2P disabled")
  }
}

func P2Pproc() {
  var start = nr.NowMillisec()
  signal.Reserve(1)
  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    p2p()
    log.Trace("p2p() thread exiting")
    signal.ExitRequest()
    nr.RecordDuration("P2Pproc() duration", start)
  }(signal.WG())
}

func P2PShutdown() {
  if ! conf.IsP2P {
    return
  }
  P2Pnode.Close()
  log.Trace("P2Pproc() terminating")
}

func p2p() {
  if ! conf.IsP2P {
    return
  }
  overlay := kademlia.New()
  if P2Pnode != nil {
    P2Pnode.Bind(overlay.Protocol())
    err := P2Pnode.Listen()
    if err != nil {
      log.Error(fmt.Sprintf("P2P listen error: %v", err))
      return
    }
  }
  log.Trace(fmt.Sprintf("P2P node listening: %v", P2Pnode.ID()))
  var  N = 0
  for ! signal.ExitRequested() {
    if N > 10 {
      ids := overlay.Table().Peers()
      for _, id := range ids {
        log.Trace(fmt.Sprintf("P2P peers %s(%s)", id.Address, id.ID.String()))
      }
      log.Trace(fmt.Sprintf("#P2P peers %d", len(ids)))
    }
    if N > conf.P2PDiscovery {
      ids := overlay.Discover()
      for _, id := range ids {
        log.Trace(fmt.Sprintf("P2P peers %s(%s)", id.Address, id.ID.String()))
      }
      log.Trace(fmt.Sprintf("#P2P peers %d", len(ids)))
    }
    stdlib.SleepForASecond()
    N++
  }
}
