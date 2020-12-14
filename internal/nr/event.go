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

func SendEvent(_payload []byte) (bool, error, *gabs.Container) {
  g, err := event(conf.Nrapi, conf.Evtapi, _payload)
  if err != nil {
    return false, err, nil
  }
  if g == nil {
    return false, nil, nil
  }
  if ! g.Exists("success") {
    return false, nil, g
  }
  val  := bool(g.Path("success").Data().(bool))
  if err != nil {
    return false, err, g
  }
  return val, nil, g
}

func event(nrikey, _url string, _payload []byte) (*gabs.Container, error) {
  var b bytes.Buffer
  url := fmt.Sprintf(_url, conf.Account)
  w := gzip.NewWriter(&b)
  w.Write([]byte(_payload))
  w.Close()
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(b.Bytes()))
  if err != nil {
    return nil, err
  }
  req.Header.Set("X-Insert-Key", nrikey)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Content-Encoding", "gzip")
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return nil, err
  } else {
    body, err := ioutil.ReadAll(resp.Body)
    if err == nil {
      r, err := gabs.ParseJSON(body)
      if err != nil {
        return nil, err
      }
      return r, nil
   } else {
      return nil, err
   }
  }
  return nil, err
}
