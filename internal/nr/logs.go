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
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
  "github.com/Jeffail/gabs"
)

func Log(msg string, logtype string, ctx logrus.Fields) {
  payload := gabs.New()
  payload.Set(conf.ID, "attributes", "applicationID")
  payload.Set(conf.Name, "attributes", "applicationName")
  payload.Set(si.SysInfo().Hostname, "hostname")
  payload.Set("tsak", "attributes", "logSource")
  for key, value := range ctx {
    payload.Set(value, "attributes", key)
  }
  payload.Set(time.Now().UnixNano() / int64(time.Millisecond), "timestamp")
  payload.Set(msg, "message")
  payload.Set(logtype, "logtype")
  if conf.Nrapi != "" {
    logs(conf.Nrapi, conf.Logapi, true, []byte(payload.String()))
  }
}

func SendLog(_payload []byte) (bool, error, *gabs.Container) {
  var payload []byte
  var b bytes.Buffer
  w := gzip.NewWriter(&b)
  w.Write([]byte(_payload))
  w.Close()
  payload = []byte(b.Bytes())
  url := conf.Logapi
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
  if err != nil {
    return false, err, nil
  }
  req.Header.Set("Api-Key", conf.Nrapi)
  req.Header.Set("Content-Type", "application/gzip")
  req.Header.Set("Content-Encoding", "gzip")
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return false, err, nil
  } else {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return false, err, nil
    } else {
      r, err := gabs.ParseJSON(body)
      if err != nil {
        return false, err, nil
      }
      if ! r.Exists("requestId") {
        return false, nil, r
      }
      return true, nil, r
    }
  }
  return false, nil, nil
}

func logs(nrikey string, url string, compress bool, _payload []byte) bool {
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
  req.Header.Set("Api-Key", nrikey)
  if compress {
    req.Header.Set("Content-Type", "application/gzip")
    req.Header.Set("Content-Encoding", "gzip")
  } else {
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Content-Encoding", "json")
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    fmt.Printf("NR Log ingestion failure: %v %v\n", resp, err)
    return false
  } else {
    // fmt.Printf("NR Log ingestion success: %v %v\n", resp, err)
    ioutil.ReadAll(resp.Body)
    telemetrydb.Counter("tsak.NR.logs")
    return true
  }
}
