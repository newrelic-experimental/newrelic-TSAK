package packages

import (
  "os"
  "bufio"
  "reflect"
  "golang.org/x/sys/unix"
  "github.com/mattn/anko/env"
)
func CanRead(f *os.File, nsec int64) bool {
  ffd := int(f.Fd())
  rFds := &unix.FdSet{}
  rFds.Set(ffd)

  n, err := unix.Select(1, rFds, nil, nil, &unix.Timeval{Sec:nsec, Usec:0})
  if err != nil {
    return false
  }
  if n > 0 {
    return true
  }
  return false
}
func init() {
  env.Packages["bufio"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(bufio.NewScanner),
    "CanRead":            reflect.ValueOf(CanRead),
  }
  env.PackageTypes["bufio"] = map[string]reflect.Type{
    "Scanner":          reflect.TypeOf(bufio.Scanner{}),
  }
}
