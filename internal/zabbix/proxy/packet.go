// Package implements Zabbix proxy protocol
package proxy

import (
	// "fmt"
	"encoding/json"
	"github.com/newrelic-experimental/newrelic-TSAK/internal/zabbix/client"
	"time"
)

// Heartbeat Packet.
type GenericPacket struct {
	Request string `json:"request"`
	Host    string `json:"host"`
	Clock   int64  `json:"clock,omitempty"`
	Version string `json:"version,omitempty"`
}

// Heartbeat Packet constructor.
func (proxy *Proxy) NewGenericPacket(request string, clock int64, _compress ...bool) *client.Packet {
	var _c bool
	ap := &GenericPacket{Request: request, Host: proxy.Name, Version: "5.2.0"}
	// use current time, if `clock` is not specified
	if ap.Clock = time.Now().Unix(); clock != 0 {
		ap.Clock = int64(clock)
	}
	if _c = true ; len(_compress) > 0 {
		_c = bool(_compress[0])
	}
	jsonData, _ := json.Marshal(ap)
	// fmt.Println("PKT",string(jsonData))
	packet := &client.Packet{Request: request, IsCompress: _c, Data: jsonData}
	return packet
}

// Availability Data.
type AvailabilityData struct {
	Hostid         int64  `json:"hostid,omitempty"`
	Available      int    `json:"available,omitempty"`
	Error          string `json:"error,omitempty"`
	Snmp_available int    `json:"snmp_available,omitempty"`
	Snmp_error     string `json:"snmp_error,omitempty"`
	Ipmi_available int    `json:"ipmi_available,omitempty"`
	Ipmi_error     string `json:"ipmi_error,omitempty"`
	Jmx_available  int    `json:"jmx_available,omitempty"`
	Jmx_error      string `json:"jmx_error,omitempty"`
}

// Host Availability Packet.
type AvailabilityPacket struct {
	Request string              `json:"request"`
	Host    string              `json:"host"`
	Data    []*AvailabilityData `json:"data,omitempty"`
	Clock   int64               `json:"clock,omitempty"`
}

// Host Availability Packet constructor.
func (proxy *Proxy) NewAvailabilityPacket(data []*AvailabilityData, clock ...int64) *client.Packet {
	ap := &AvailabilityPacket{Request: `host availability`, Host: proxy.Name, Data: data}
	// use current time, if `clock` is not specified
	if ap.Clock = time.Now().Unix(); len(clock) > 0 {
		ap.Clock = int64(clock[0])
	}
	jsonData, _ := json.Marshal(ap)
	packet := &client.Packet{Request: `host availability`, Data: jsonData}
	return packet
}

// History Data.
type HistoryData struct {
	Host  string      `json:"host,omitempty"`
	Key   string      `json:"key,omitempty"`
	Clock int64       `json:"clock,omitempty"`
	Ns    int64       `json:"ns,omitempty"`
	State string      `json:"state,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// Host History Packet.
type HistoryPacket struct {
	Request string         `json:"request"`
	Host    string         `json:"host"`
	Data    []*HistoryData `json:"data,omitempty"`
	Clock   int64          `json:"clock,omitempty"`
}

// Host History Packet constructor.
func (proxy *Proxy) NewHistoryPacket(data []*HistoryData, clock ...int64) *client.Packet {
	ap := &HistoryPacket{Request: `history data`, Host: proxy.Name, Data: data}
	// use current time, if `clock` is not specified
	if ap.Clock = time.Now().Unix(); len(clock) > 0 {
		ap.Clock = int64(clock[0])
	}
	jsonData, _ := json.Marshal(ap)
	packet := &client.Packet{Request: `history data`, Data: jsonData}
	return packet
}
