package zabbix

import (
  "fmt"
  "regexp"
  "strings"
  "encoding/csv"
  glob "github.com/ganbarodigital/go_glob"
  // "github.com/hjson/hjson-go"
  // "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)



func ParseParamsInZabbixKey(key string, params string) (args map[string]string, err error) {
  var found bool
  var glb  *glob.Glob

  args = make(map[string]string)
  csvr := csv.NewReader(strings.NewReader(params))
  record, err := csvr.Read()
  if err != nil {
    fmt.Println("Ouch", err)
    return
  }
  found = false
  for _, zv := range(ZKEYS) {
    patt, ok := zv.(map[string]interface{})["pattern"]
    if ! ok {
      continue
    }
    glb = glob.NewGlob(patt.(string))
    isMatch, _ := glb.Match(key)
    if isMatch {
      found = true
      fields := zv.(map[string]interface{})["keys"].([]interface{})
      for n, v := range(record) {
        if n < len(fields) {
          args[fields[n].(string)] = v
        } else {
          args[fmt.Sprintf("ARG%v", n)] = v
        }
      }
      break
    }
  }
  if ! found {
    for n, v := range(record) {
      k := fmt.Sprintf("ARG%v", n)
      args[k] = v
    }
  }
  return
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
    args, _ = ParseParamsInZabbixKey(key, v[0][2])
  }
  return
}
