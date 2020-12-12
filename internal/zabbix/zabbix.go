package zabbix

import (
  "fmt"
  "bytes"
  "net"
  "time"
  "regexp"
  "strings"
  "compress/zlib"
  "encoding/binary"
  "io/ioutil"
  "github.com/Jeffail/gabs"
  "github.com/sirupsen/logrus"
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

func MakeData(req string, data string, host string, compress bool) []byte {
  return MakePacket(fmt.Sprintf(`{"request" : "%s", "host" : "%s", "data" : [%s]}`, req, host, data), compress)
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

func ParseKey(key string) (name string, args map[string]string) {
  args = make(map[string]string)
  name = ""
  re := regexp.MustCompile(`^(.+)\[(.*)\]$`)
  if ! re.MatchString(key) {
    name = key
    return
  }
  v := re.FindAllStringSubmatch(key, 2)
  name = v[0][1]
  if len(strings.TrimSpace(v[0][2])) > 0 {
    for n, s := range strings.Split(v[0][2], ",") {
      args[fmt.Sprintf("ARGS%v",n)] = strings.TrimSpace(s)
    }
  }
  return
}

func ParseRaw(data []byte) []byte {
  if len(data) < 13 {
    log.Trace("ZABBIX packet is malformed")
    return nil
  }
  payload := data[13:]
  return payload
}

func OneWay(dst string, zpkt []byte, timeout uint64) []byte {
  c, err := net.DialTimeout("tcp", dst, time.Second*time.Duration(timeout))
  defer c.Close()
  if err == nil {
    c.Write(zpkt)
    res, err := ioutil.ReadAll(c)
    if err == nil {
      return res
    } else {
      log.Trace("ZABBIX 1-way read error", logrus.Fields{
        "error": err,
        "destination": dst,
      })
    }
  } else {
    log.Trace("ZABBIX 1-way network error", logrus.Fields{
      "error": err,
      "destination": dst,
    })
  }
  return nil
}

func TwoWay(dst string, zpkt []byte, timeout uint64) bool {
  var pkt *gabs.Container
  c, err := net.DialTimeout("tcp", dst, time.Second*time.Duration(timeout))
  defer c.Close()
  if err == nil {
    c.Write(zpkt)
    res, err := ioutil.ReadAll(c)
    pkt = nil
    if err == nil {
      pkt = Parse(res)
    } else {
      log.Trace("ZABBIX 3-way reading error", logrus.Fields{"error":err, "destination":dst})
      return false
    }
    if pkt == nil {
      log.Trace("ZABBIX 3-way JSON parsing error", logrus.Fields{"error":err, "destination":dst})
      return false
    }
    return "\"success\"" == pkt.Search("response").String()
  } else {
    log.Trace("ZABBIX 3-way network error", logrus.Fields{"error":err, "destination":dst})
  }
  return false
}

func ThreeWay(dst string, zpkt []byte, zconfirmation []byte, timeout uint64) *gabs.Container {
  var pkt *gabs.Container
  c, err := net.DialTimeout("tcp", dst, time.Second*time.Duration(timeout))
  defer c.Close()
  if err == nil {
    c.Write(zpkt)
    res, err := ioutil.ReadAll(c)
    pkt = nil
    if err == nil {
      pkt = Parse(res)
      c.Write(zconfirmation)
    } else {
      log.Trace("ZABBIX 3-way parsing error", logrus.Fields{"error":err, "destination":dst})
    }
    return pkt
  } else {
    log.Trace("ZABBIX 3-way network error", logrus.Fields{"error":err, "destination":dst})
  }
  return nil
}
