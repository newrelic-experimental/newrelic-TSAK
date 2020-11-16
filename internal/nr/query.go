package nr

import (
  "fmt"
  "net/http"
  "net/url"
  "io/ioutil"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/Jeffail/gabs"
)

func Query(_query string) []*gabs.Container {
  if conf.Nrapiq != "" {
    res := query(conf.Nrapiq, conf.Queryapi, _query)
    ret, err := gabs.ParseJSON([]byte(res))
    if err != nil {
      return nil
    }
    cont := processQueryResult(ret)
    return cont
  } else {
    return nil
  }
}

func processQueryResult(r *gabs.Container) []*gabs.Container {
  var cont []*gabs.Container
  res1, err := r.S("results").Children()
  if err != nil {
    return cont
  }
  for _, c := range res1 {
    res2, err := c.S("events").Children()
    if err != nil {
      continue
    }
    for _, c1 := range res2 {
      cont = append(cont, c1)
    }
  }
  return cont
}

func query(nriqkey string, _url string, query string) string {
  url := fmt.Sprintf(_url, conf.Account, url.QueryEscape(query))
  fmt.Println(url)
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return ""
  }
  req.Header.Set("X-Query-Key", nriqkey)
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return ""
  } else {
    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return ""
    }
    return string(res)
  }
}
