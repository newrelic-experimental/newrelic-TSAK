package packages

import (
  "context"
	"fmt"
	"net"
	"time"
  "crypto/tls"
  "crypto/x509"
	"github.com/pion/dtls/v2"
  "github.com/pion/dtls/v2/examples/util"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "reflect"
  "github.com/mattn/anko/env"
)

func UDPDTLSNew(bindip string, port int, srvCert, srvPub string) net.Listener {
  addr := &net.UDPAddr{IP: net.ParseIP(bindip), Port: port}
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  certificate, err := util.LoadKeyAndCertificate(srvCert, srvPub)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  rootCertificate, err := util.LoadCertificate(srvPub)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  certPool := x509.NewCertPool()
  cert, err := x509.ParseCertificate(rootCertificate.Certificate[0])
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  certPool.AddCert(cert)
  config := &dtls.Config{
		Certificates:         []tls.Certificate{*certificate},
		ExtendedMasterSecret: dtls.RequireExtendedMasterSecret,
		ClientAuth:           dtls.RequireAndVerifyClientCert,
		ClientCAs:            certPool,
		ConnectContextMaker: func() (context.Context, func()) {
			return context.WithTimeout(ctx, 30*time.Second)
		},
	}
  listener, err := dtls.Listen("udp", addr, config)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  return listener
}

func UDPDTLSDial(dialip string, port int, cliCert, cliPub string) *dtls.Conn {
  addr := &net.UDPAddr{IP: net.ParseIP(dialip), Port: port}
  certificate, err := util.LoadKeyAndCertificate(cliCert, cliPub)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  rootCertificate, err := util.LoadCertificate(cliPub)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  certPool := x509.NewCertPool()
  cert, err := x509.ParseCertificate(rootCertificate.Certificate[0])
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  certPool.AddCert(cert)
  config := &dtls.Config{
		Certificates:         []tls.Certificate{*certificate},
		ExtendedMasterSecret: dtls.RequireExtendedMasterSecret,
		RootCAs:              certPool,
	}
  dtlsConn, err := dtls.Dial("udp", addr, config)
  if err != nil {
    log.Trace(fmt.Sprintf("DTLS.error: %v", err))
  }
  return dtlsConn
}

func init() {
  env.Packages["net/dtls"] = map[string]reflect.Value{
    "New":    reflect.ValueOf(UDPDTLSNew),
    "Dial":   reflect.ValueOf(UDPDTLSDial),
  }
  env.PackageTypes["net/dtls"] = map[string]reflect.Type{
    "Conn":           reflect.TypeOf(dtls.Conn{}),

  }
}
