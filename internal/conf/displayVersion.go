package conf

import (
  "os"
  "fmt"
  "github.com/olekukonko/tablewriter"
)

func DisplayVersion() {
  data  := [][]string{
    []string{"TSAK version", Ver},
    []string{"TSAK version major", fmt.Sprintf("%v",VerMaj)},
    []string{"TSAK version minor", fmt.Sprintf("%v",VerMin)},
    []string{"TSAK version prerelease", fmt.Sprintf("%v",VerPrerelease)},
  }
  table := tablewriter.NewWriter(os.Stdout)
  for _, v := range data {
    table.Append(v)
  }
  table.Render()
}
