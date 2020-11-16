package nr

import (
  "fmt"
  "bytes"
  "compress/gzip"
  "net/http"
  "time"
  "io/ioutil"
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/si"
  "github.com/Jeffail/gabs"
)


func Event(evttype string, ctx logrus.Fields) {
  payload := gabs.New()
  payload.Set(evttype, "eventType")
  payload.Set(conf.ID,  "applicationID")
  payload.Set(conf.Name,  "applicationName")
  payload.Set(si.SysInfo().Hostname, "hostname")
  payload.Set("tsak", "evtSource")
  for key, value := range ctx {
    payload.Set(value, key)
  }
  payload.Set(time.Now().UnixNano() / int64(time.Millisecond), "timestamp")
  if conf.Nrapi != "" {
    event(conf.Nrapi, conf.Evtapi, []byte(payload.String()))
  }
}

func SendEvent(_payload []byte) bool {
  return event(conf.Nrapi, conf.Evtapi, _payload)
}

func event(nrikey, _url string, _payload []byte) bool {
  var b bytes.Buffer
  url := fmt.Sprintf(_url, conf.Account)
  w := gzip.NewWriter(&b)
  w.Write([]byte(_payload))
  w.Close()
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(b.Bytes()))
  if err != nil {
    return false
  }
  req.Header.Set("X-Insert-Key", nrikey)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Content-Encoding", "gzip")
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return false
  } else {
    ioutil.ReadAll(resp.Body)
  }
  return true
}
