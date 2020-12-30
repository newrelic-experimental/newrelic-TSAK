package conf

import (
  "strings"
)

func ParseArgs() (args []string, res map[string]interface{}) {
  var prev=""
  res = make(map[string]interface{})
  for _,c :=  range(Args) {
    if strings.HasPrefix(c, "/") && prev == "" {
      prev = c[1:]
    } else if ! strings.HasPrefix(c, "/") && prev != "" {
      res[strings.ToUpper(prev)] = c
      prev = ""
    } else if strings.HasPrefix(c, "/") && prev != "" {
      res[strings.ToUpper(prev)] = true
      prev = c[1:]
    } else if ! strings.HasPrefix(c, "/") && prev == "" {
      args = append(args, c)
    } else {
      args = append(args, c)
    }
  }
  if prev != "" {
    res[strings.ToUpper(prev)] = true
  }
  return
}
