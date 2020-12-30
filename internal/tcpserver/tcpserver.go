package tcpserver

import (
  "fmt"
  "time"
	"io"
  "bufio"
	"crypto/tls"
	"net"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

// Client holds info about connection
type Client struct {
	conn     net.Conn
  Bufsize  int64
  MsgSize  int64
  IsText   bool
	Server   *AServer
}

// TCP server
type AServer struct {
	address                  string // Address to open connection: localhost:9999
	config                   *tls.Config
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message []byte)
  onMessageHeader          func(c *Client, message []byte)
}

// Read client data from channel
func (c *Client) listen() {
  var message []byte
  var n int
  var err error
  message = make([]byte, c.Bufsize)
  log.Trace(fmt.Sprintf("Establishing connection from %v to %v", c.conn.RemoteAddr().String(), c.conn.LocalAddr().String()))
	c.Server.onNewClientCallback(c)
  reader := bufio.NewReader(c.conn)
	for ! signal.ExitRequested() {
    if c.IsText {
      msg, e := reader.ReadString('\n')
      message = []byte(msg)
      n = len(message)
      err = e
    } else {
      n, err = io.ReadAtLeast(reader, message, int(c.Bufsize))
    }
		if err != nil || n == 0{
      log.Trace(fmt.Sprintf("Closing connection from %v to %v", c.conn.RemoteAddr().String(),c.conn.LocalAddr().String()))
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, message)
	}
}

// Send text message to client
func (c *Client) Send(message string) error {
	_, err := c.conn.Write([]byte(message))
	return err
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *AServer) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *AServer) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}


// Called when Client receives new message
func (s *AServer) OnNewMessage(callback func(c *Client, message []byte)) {
	s.onNewMessage = callback
}

// Listen starts network server
func (s *AServer) Listen(n int64, bufsize int64, istxt bool) {
	var listener net.Listener
	var err error
	if s.config == nil {
		listener, err = net.Listen("tcp", s.address)
	} else {
		listener, err = tls.Listen("tcp", s.address, s.config)
	}
	if err != nil {
		log.Error("Error starting TCP server.")
	}

  go func(listener net.Listener) {
    log.Trace(fmt.Sprintf("Entering listening loop for %v", s.address))
    for {
      conn, _ := listener.Accept()
      if signal.ExitRequested() {
        break
      }
  		client := &Client{
  			conn:   conn,
        Bufsize: bufsize,
        IsText: istxt,
  			Server: s,
  		}
  		go client.listen()
    }
  }(listener)
	for ! signal.ExitRequested() {
		time.Sleep(time.Duration(n)*time.Second)
	}
  log.Trace(fmt.Sprintf("Exit requested for Loop(%v)", s.address))
  listener.Close()
}

// Creates new tcp server instance
func New(address string) *AServer {
	log.Trace(fmt.Sprintf("Creating non-encrypted server with address %v", address))
	server := &AServer{
		address: address,
		config:  nil,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message []byte) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}

func NewWithTLS(address string, certFile string, keyFile string) *AServer {
	log.Trace(fmt.Sprintf("Creating TLS server with address %v", address))
	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	server := &AServer{
		address: address,
		config:  &config,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message []byte) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}
