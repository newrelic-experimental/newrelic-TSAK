package packages

import (
  "github.com/dspinhirne/netaddr-go"
  "reflect"
  "github.com/mattn/anko/env"
)

func NetList(net *netaddr.IPv4Net) []string {
  res := []string{}
  networkAddr := net.Network()
  ip := networkAddr.Next()
  for c := 1; c < int(net.Len()-1); c++ {
    res = append(res, ip.String())
    ip = ip.Next()
  }
  return res
}

func init() {
  env.Packages["netaddr"] = map[string]reflect.Value{
    "IP":             reflect.ValueOf(netaddr.ParseIP),
    "IP4":            reflect.ValueOf(netaddr.ParseIPv4),
    "Net":            reflect.ValueOf(netaddr.ParseIPNet),
    "Range":          reflect.ValueOf(NetList),
  }
  env.PackageTypes["netaddr"] = map[string]reflect.Type{
    "IPv4":           reflect.TypeOf(netaddr.IPv4{}),
  }
}
