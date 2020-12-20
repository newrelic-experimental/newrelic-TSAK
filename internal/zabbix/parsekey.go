package zabbix

import (
  "fmt"
  "regexp"
  "strings"
  "encoding/csv"
)

var ZABBIXKEYS = `
{

}
`

func ParseParamsInZabbixKey(params string) (args map[string]string, err error) {
  var found bool
  args = make(map[string]string)
  csvr := csv.NewReader(strings.NewReader(params))
  record, err := csvr.Read()
  if err != nil {
    fmt.Println("Ouch", err)
    return
  }
  found = false
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
    args, _ = ParseParamsInZabbixKey(v[0][2])
  }
  return
}
