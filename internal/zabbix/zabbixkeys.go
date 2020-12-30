package zabbix

import (
  "github.com/hjson/hjson-go"
)

var ZKEYS map[string]interface{}


var ZABBIXKEYS = `
{
  log: {
    pattern: log\[*\]
    keys: [
      file
      regexp
      encoding
      maxlines
      mode
      output
      maxdelay
      options
    ]
  }
  logCount: {
    pattern: log.count\[*\]
    keys: [
      file
      regexp
      maxproxlines
      mode
      maxdelay
      options
    ]
  }
  modBusGet: {
    pattern: modbus.get*
    keys: [
      endpoint
      slave_id
      function
      address
      colunt
      type
      endianness
      offset
    ]
  }
  netDNS: {
    pattern: net.dns*
    keys: [
      ip
      name
      type
      timeout
      count
      protocol
    ]
  }
  netIfCollision: {
    pattern: net.if.collisions*
    keys: [
      if
    ]
  }
  netIf: {
    pattern: net.if.*\[*\]
    keys: [
      if
      mode
    ]
  }
  netSomeListen: {
    pattern: net.*.listen*
    keys: [
      port
    ]
  }
  netTcpPort: {
    pattern: net.tcp.port*
    keys: [
      ip
      port
    ]
  }
  netSomeService: {
    pattern: net.*.service*
    keys: [
      service
      ip
      port
    ]
  }
  procCpuUtil: {
    pattern: proc.cpu.util*
    keys: [
      name
      user
      type
      cmdline
      mode
      zone
    ]
  }
  procMem: {
    pattern: proc.mem*
    keys: [
      name
      user
      mode
      cmdline
      memtype
    ]
  }
  procNum: {
    pattern: proc.num*
    keys: [
      name
      user
      state
      cmdline
      zone
    ]
  }
  sensor: {
    pattern: sensor\[*\]
    keys: [
      device
      sensor
      mode
    ]
  }
  scu: {
    pattern: system.cpu.util*
    keys: [
      cpu
      type
      mode
      logicalOrphysical
    ]
  }
  shn: {
    pattern: system.hostname*
    keys: [
      type
    ]
  }
  shcpu: {
    pattern: system.hw.cpu*
    keys: [
      cpu
      info
    ]
  }
  shch: {
    pattern: system.hw.chassis*
    keys: [
      info
    ]
  }
  shdev: {
    pattern: system.hw.devices*
    keys: [
      type
    ]
  }
  shmac: {
    pattern: system.hw.macaddr*
    keys: [
      interface
      format
    ]
  }
  sltime: {
    pattern: system.localtime*
    keys: [
      type
    ]
  }
  srun: {
    pattern: system.run*
    keys: [
      command
      mode
    ]
  }
  sstat: {
    pattern: system.stat*
    keys: [
      resource
      type
    ]
  }
  sswos: {
    pattern: system.sw.os*
    keys: [
      info
    ]
  }
  sswpkg: {
    pattern: system.sw.packages*
    keys: [
      package
      manager
      format
    ]
  }
  sswap: {
    pattern: system.swap.*\[*\]
    keys: [
      device
      type
    ]
  }
  vfsdev: {
    pattern: vfs.dev.*\[*\]
    keys: [
      device
      type
      mode
    ]
  }
  vfsdircount: {
    pattern: vfs.dir.count*
    keys: [
      dir
      regexincl
      regexexcl
      typesincl
      typesexcl
      maxdepth
      minsize
      maxsize
      minage
      maxage
      regexexcldir
    ]
  }
  vfsdirsize: {
    pattern: vfs.dir.size*
    keys: [
      dir
      regexincl
      regexexcl
      mode
      maxdepth
      regexexcldir
    ]
  }
  vfsfilechk: {
    pattern: vfs.file.cksum*
    keys: [
      file
    ]
  }
  vfsfilcnt: {
    pattern: vfs.file.contents*
    keys: [
      file
      encoding
    ]
  }
  vfsfilex: {
    pattern: vfs.file.exists*
    keys: [
      file
      typesincl
      typesexcl
    ]
  }
  vfsfilechkmd5: {
    pattern: vfs.file.md5*
    keys: [
      file
    ]
  }
  vfsfilereg: {
    pattern: vfs.file.regexp*
    keys: [
      file
      regexp
      encoding
      startline
      endline
      output
    ]
  }
  vfsfileregm: {
    pattern: vfs.file.regmatch*
    keys: [
      file
      regexp
      encoding
      startline
      endline
    ]
  }
  vfsfiletime: {
    pattern: vfs.file.time*
    keys: [
      file
      mode
    ]
  }
  vfsfilsize: {
    pattern: vfs.file.size*
    keys: [
      file
    ]
  }
  vfsfs: {
    pattern: vfs.fs*\[*\]
    keys: [
      fs
      mode
    ]
  }
  vfsmemsize: {
    pattern: vfs.memory.size*
    keys: [
      mode
    ]
  }
  wpg: {
    pattern: web.page.get*
    keys: [
      host
      path
      port
    ]
  }
  wpp: {
    pattern: web.page.perf*
    keys: [
      host
      path
      port
    ]
  }
  wpregex: {
    pattern: web.page.regexp*
    keys: [
      host
      path
      port
      regexp
      length
      output
    ]
  }
  name: {
    pattern: ....
    keys: [

    ]
  }
}
`

func init() {
  hjson.Unmarshal([]byte(ZABBIXKEYS), &ZKEYS)
}
