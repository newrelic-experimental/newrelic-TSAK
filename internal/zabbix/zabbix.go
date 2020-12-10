package zabbix

import (
  "bytes"
  "compress/zlib"
  "encoding/binary"
  "io/ioutil"
  "github.com/Jeffail/gabs"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

func uncompressDataLen(data []byte) []byte {
	dataLen := make([]byte, 8)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(data)))
	return dataLen
}

func decompressJSON(data []byte) []byte {
	buf := bytes.NewReader(data)
	r, err := zlib.NewReader(buf)
	if err != nil {
    log.Trace("Error uncompressing ZABBIX JSON payload")
		return nil
	}
	ndata, err := ioutil.ReadAll(r)
	if err != nil {
    log.Trace("Error reading ZABBIX JSON payload")
		return nil
	}
	return ndata
}

func getHeader() []byte {
	return []byte("ZBXD\x01")
}
func getHeaderCompress() []byte {
	return []byte("ZBXD\x03")
}

func anyByteCompress(data []byte) (cdata []byte, dataLen1 []byte, dataLen2 []byte) {
	var buf bytes.Buffer
	dataLen1 =  make([]byte, 4)
	dataLen2 =  make([]byte, 4)
	binary.LittleEndian.PutUint32(dataLen1, uint32(len(data)))
	w := zlib.NewWriter(&buf)
	w.Write(data)
	w.Close()
	cdata = buf.Bytes()
	binary.LittleEndian.PutUint32(dataLen2, uint32(len(cdata)))
	return
}

func MakePacket(data string, isCompress bool) []byte {
  var buffer []byte

  if isCompress {
    cdata, dlorig, dlcomp := anyByteCompress([]byte(data))
    buffer = append(getHeaderCompress(), dlcomp...)
    buffer = append(buffer, dlorig...)
    buffer = append(buffer, cdata...)
  } else {
    buffer = append(getHeader(), uncompressDataLen([]byte(data))...)
    buffer = append(buffer, data...)
  }
  return buffer
}

func MakeResp(resp string, version string, compress bool) []byte {
  pkt := gabs.New()
  pkt.Set(resp, "response")
  pkt.Set(version, "version")
  return MakePacket(pkt.String(), compress)
}

func MakeReq(resp string, host string, compress bool) []byte {
  pkt := gabs.New()
  pkt.Set(resp, "request")
  pkt.Set(host, "host")
  return MakePacket(pkt.String(), compress)
}

func MakeData() {
  
}

func GetPayloadSize(header []byte) (uint32, bool) {
  if string(header[:4]) != "ZBXD" {
    log.Trace("ZABBIX packet is missing it's signature")
    return 0, false
  }
  if header[4] == 1 {
    return binary.LittleEndian.Uint32(header[5:13]), false
  } else if header[4] == 3 {
    return binary.LittleEndian.Uint32(header[5:9]), true
  } else {
    log.Trace("Unknown flag in ZABBIX header")
    return 0, false
  }
}

func Parse(data []byte) *gabs.Container {
  if len(data) < 13 {
    log.Trace("ZABBIX packet is malformed")
    return nil
  }
  hdr := data[:13]
  payload := data[13:]
  return ParsePacket(hdr, payload)
}

func ParsePacket(header, data []byte) *gabs.Container {
  size, compress := GetPayloadSize(header)
  if size == 0 {
    log.Trace("Empty payload in ZABBIX packet")
    return nil
  }
  if compress {
    pkt := decompressJSON(data)
    if pkt == nil {
      return nil
    }
    npkt, err := gabs.ParseJSON(pkt)
    if err == nil {
      return npkt
    } else {
      return nil
    }
  } else {
    pkt, err := gabs.ParseJSON(data)
    if err == nil {
      return pkt
    } else {
      return nil
    }
  }
}
