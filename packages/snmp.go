package packages

import (
  "net"
  "time"
  "github.com/deejross/go-snmplib"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/snmp"
  "reflect"
  "github.com/mattn/anko/env"
)

func UnmarshalTrap(sock *net.IPConn, b []byte, n int) snmplib.Trap {
  snmp := snmplib.NewSNMPOnConn("", "", snmplib.SNMPv2c, 2*time.Second, 5, sock)
  msg := b[:n]
  varbinds, _ := snmp.ParseTrap(msg)
  return varbinds
}

func UnmarshalTrap3(users []snmplib.V3user, sock *net.IPConn, b []byte, n int) snmplib.Trap {
  snmp := snmplib.NewSNMPOnConn("", "", snmplib.SNMPv3, 2*time.Second, 5, sock)
  msg := b[:n]
  snmp.TrapUsers = users
  varbinds, _ := snmp.ParseTrap(msg)
  return varbinds
}


func init() {
  env.Packages["protocols/snmp"] = map[string]reflect.Value{
    "ParseTrap":  reflect.ValueOf(UnmarshalTrap),
    "ParseTrap3": reflect.ValueOf(UnmarshalTrap3),
    "InitMib":    reflect.ValueOf(snmp.InitMib),
    "LoadModule": reflect.ValueOf(snmp.LoadModule),
    "LoadAll":    reflect.ValueOf(snmp.InitAndLoadAll),
    "OID":        reflect.ValueOf(snmp.OID),
    "IsOID":      reflect.ValueOf(snmp.IsOID),
    "SYMBOL":     reflect.ValueOf(snmp.SYMBOL),
    "Init":       reflect.ValueOf(snmp.InitSNMP),
    "Client":     reflect.ValueOf(snmplib.NewSNMP),
    "SNMPv1":     reflect.ValueOf(snmplib.SNMPv1),
    "SNMPv2c":    reflect.ValueOf(snmplib.SNMPv2c),
    "SNMPv3":     reflect.ValueOf(snmplib.SNMPv3),
    "ParseOID":   reflect.ValueOf(snmplib.MustParseOid),
  }
  env.PackageTypes["protocols/snmp"] = map[string]reflect.Type{
    "SNMP":           reflect.TypeOf(snmplib.SNMP{}),
    "SNMPConf":       reflect.TypeOf(snmp.SNMPConf{}),
    "V3user":         reflect.TypeOf(snmplib.V3user{}),
  }
}
