package nr

import (
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

func Metric(mname string, _type string, value interface{}, ctx logrus.Fields) {
  payload := gabs.New()
  payload.Set(mname, "name")
  payload.Set(value, "value")
  payload.Set(_type, "type")
  payload.Set(conf.ID,  "attributes", "applicationID")
  payload.Set(conf.Name,  "attributes", "applicationName")
  payload.Set(si.SysInfo().Hostname, "attrubutes", "hostname")
  payload.Set("tsak", "attributes", "metricSource")
  for key, value := range ctx {
    payload.Set(value, "attributes", key)
  }
  payload.Set(time.Now().UnixNano() / int64(time.Millisecond), "timestamp")
  if conf.Nrapi != "" {
    metrics(conf.Nrapi, conf.Metricapi, true, []byte(payload.String()))
  }
}

func SendMetric(_payload []byte) bool {
  return metrics(conf.Nrapi, conf.Metricapi, true, _payload)
}

func metrics(nrikey string, url string, compress bool, _payload []byte) bool {
  var payload []byte
  var b bytes.Buffer
  if compress {
    w := gzip.NewWriter(&b)
    w.Write([]byte(_payload))
    w.Close()
    payload = []byte(b.Bytes())
  } else {
    payload = []byte(_payload)
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
  if err != nil {
    return false
  }
  req.Header.Set("X-Insert-Key", nrikey)
  req.Header.Set("Content-Type", "application/json")
  if compress {
    req.Header.Set("Content-Encoding", "gzip")
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return false
  }
  ioutil.ReadAll(resp.Body)
  return true
}
