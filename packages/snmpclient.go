package packages

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/k-sone/snmpgo"
  "reflect"
  "github.com/mattn/anko/env"
)

func SNMPv1Get(addr string, community string, _oid string, retry uint) interface{} {
  snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V1,
		Address:   addr,
		Retries:   retry,
		Community: community,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error in SNMPGet init %s", err))
		return false
	}
  oids, err := snmpgo.NewOids([]string{
    _oid,
  })
  if err != nil {
		// Failed to parse Oids
		log.Error(fmt.Sprintf("Error in SNMPGet OID init %s", err))
		return false
	}

	if err = snmp.Open(); err != nil {
		log.Error(fmt.Sprintf("SNMPGet failed to open connection %s", err))
		return false
	}
	defer snmp.Close()

	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		log.Error(fmt.Sprintf("SNMPGet is failed to request %s", err))
		return false
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		log.Warning(fmt.Sprintf("SNMPGet failed at agent: %v %v", pdu.ErrorStatus(), pdu.ErrorIndex()))
	}

	res := pdu.VarBinds().MatchOid(oids[0])
  return res.Variable
}

func SNMPv2cGet(addr string, community string, _oid string, retry uint) interface{} {
  snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   addr,
		Retries:   retry,
		Community: community,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error in SNMPGet init %s", err))
		return false
	}
  oids, err := snmpgo.NewOids([]string{
    _oid,
  })
  if err != nil {
		// Failed to parse Oids
		log.Error(fmt.Sprintf("Error in SNMPGet OID init %s", err))
		return false
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		log.Error(fmt.Sprintf("SNMPGet failed to open connection %s", err))
		return false
	}
	defer snmp.Close()

	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		log.Error(fmt.Sprintf("SNMPGet is failed to request %s", err))
		return false
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		log.Warning(fmt.Sprintf("SNMPGet failed at agent: %v %v", pdu.ErrorStatus(), pdu.ErrorIndex()))
	}

	res := pdu.VarBinds().MatchOid(oids[0])
  return res.Variable
}

func SNMPv2cWalk(addr string, community string, _oid string, retry uint) snmpgo.VarBinds {
  var nonRepeaters = 0
  var maxRepetitions = 10
  snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   addr,
		Retries:   retry,
		Community: community,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error in SNMPWalk init %s", err))
		return nil
	}
  oids, err := snmpgo.NewOids([]string{
    _oid,
  })
  if err != nil {
		// Failed to parse Oids
		log.Error(fmt.Sprintf("Error in SNMPWalk OID init %s", err))
		return nil
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		log.Error(fmt.Sprintf("SNMPWalk is failed to request %s", err))
		return nil
	}
	defer snmp.Close()
  pdu, err := snmp.GetBulkWalk(oids, nonRepeaters, maxRepetitions)
	if err != nil {
		log.Error(fmt.Sprintf("SNMPWalk is failed to request %s", err))
		return nil
	}

	if pdu.ErrorStatus() != snmpgo.NoError {
		log.Warning(fmt.Sprintf("SNMPGet failed at agent: %v %v", pdu.ErrorStatus(), pdu.ErrorIndex()))
	}

  res := pdu.VarBinds()
  return res
}

func init() {
  env.Packages["snmp/client"] = map[string]reflect.Value{
    "Getv1":     reflect.ValueOf(SNMPv1Get),
    "Getv2c":    reflect.ValueOf(SNMPv2cGet),
    "Walk":      reflect.ValueOf(SNMPv2cWalk),
  }
  env.PackageTypes["snmp/client"] = map[string]reflect.Type{
    "VarBinds":          reflect.TypeOf(snmpgo.VarBind{}),
  }
}
