package packages

import (
  "os"
  "fmt"
  "net"
  "time"
  "path"
  "strings"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/deejross/go-snmplib"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/snmp"
  // "github.com/danwakefield/fnmatch"
  dw "github.com/karrick/godirwalk"
  "reflect"
  "github.com/mattn/anko/env"
)

func UnmarshalTrap(sock *net.IPConn, b []byte, n int) snmplib.Trap {
  snmp := snmplib.NewSNMPOnConn("", "", snmplib.SNMPv2c, 2*time.Second, 5, sock)
  msg := b[:n]
  varbinds, _ := snmp.ParseTrap(msg)
  return varbinds
}

func InitAndLoadAll(mibdirpath string) int {
  log.Trace(fmt.Sprintf("Loading and initializing MIB modules from: %s", mibdirpath))
  c := 0
  snmp.InitMib(mibdirpath)
  err := dw.Walk(mibdirpath, &dw.Options{
        Callback: func(osPathname string, de *dw.Dirent) error {
            fileStat, err := os.Stat(osPathname)
            if err != nil {
              return err
            }
            if fileStat.IsDir() {
              return nil
            }
            fn := path.Base(path.Clean(osPathname))
            fn = strings.TrimSuffix(fn, path.Ext(fn))
            log.Trace(fmt.Sprintf("Loading %s", fn))
            snmp.LoadModule(fn)
            c+=1
            return nil
        },
        Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
  })
  if err != nil {
    log.Error(fmt.Sprintf("Error scanning MIB tree: %s", err))
  }
  log.Trace(fmt.Sprintf("%d modules been loaded", c))
  return c
}

func init() {
  env.Packages["snmp"] = map[string]reflect.Value{
    "ParseTrap":  reflect.ValueOf(UnmarshalTrap),
    "InitMib":    reflect.ValueOf(snmp.InitMib),
    "LoadModule": reflect.ValueOf(snmp.LoadModule),
    "LoadAll":    reflect.ValueOf(InitAndLoadAll),
    "OID":        reflect.ValueOf(snmp.OID),
    "SYMBOL":     reflect.ValueOf(snmp.SYMBOL),
    "Client":     reflect.ValueOf(snmplib.NewSNMP),
    "SNMPv1":     reflect.ValueOf(snmplib.SNMPv1),
    "SNMPv2c":    reflect.ValueOf(snmplib.SNMPv2c),
    "SNMPv3":     reflect.ValueOf(snmplib.SNMPv3),
    "ParseOID":   reflect.ValueOf(snmplib.MustParseOid),
  }
  env.PackageTypes["snmp"] = map[string]reflect.Type{
    "SNMP":          reflect.TypeOf(snmplib.SNMP{}),
  }
}
