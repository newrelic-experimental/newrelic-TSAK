package conf

import (
  "os"
  "fmt"
  "github.com/olekukonko/tablewriter"
  "github.com/shirou/gopsutil/v3/host"
)

func DisplayVersion() {
  var data [][]string
  info, err  := host.Info()
  if err != nil {
    data  = [][]string{
      []string{"TSAK version", Ver},
      []string{"TSAK version major", fmt.Sprintf("%v",VerMaj)},
      []string{"TSAK version minor", fmt.Sprintf("%v",VerMin)},
      []string{"TSAK version prerelease", fmt.Sprintf("%v",VerPrerelease)},
      []string{"CLIPS version", "6.31"},
    }
  } else {
    data  = [][]string{
      []string{"Application ID", ID},
      []string{"TSAK version", Ver},
      []string{"TSAK version major", fmt.Sprintf("%v",VerMaj)},
      []string{"TSAK version minor", fmt.Sprintf("%v",VerMin)},
      []string{"TSAK version prerelease", fmt.Sprintf("%v",VerPrerelease)},
      []string{"CLIPS version", "6.31"},
      []string{"Host OS",           info.OS},
      []string{"Platform",          info.Platform},
      []string{"Platform family",   info.PlatformFamily},
      []string{"Platform version",  info.PlatformVersion},
      []string{"Kernel version",    info.KernelVersion},
      []string{"Architecture",      info.KernelArch},
      []string{"PID file",          fmt.Sprintf("%s/%s.pid", AppPath, ID)},

    }
  }
  table := tablewriter.NewWriter(os.Stdout)
  table.SetAlignment(tablewriter.ALIGN_LEFT)
  for _, v := range data {
    table.Append(v)
  }
  table.Render()
}
