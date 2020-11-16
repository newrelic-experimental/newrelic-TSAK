package si

import (
  "github.com/elastic/go-sysinfo"
  "github.com/elastic/go-sysinfo/types"
)

var host,_ = sysinfo.Host()

func SysInfo() types.HostInfo {
  return host.Info()
}
