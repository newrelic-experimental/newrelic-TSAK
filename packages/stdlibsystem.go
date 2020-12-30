package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/shirou/gopsutil/v3/mem"
  "github.com/shirou/gopsutil/v3/cpu"
  "github.com/shirou/gopsutil/v3/disk"
  "github.com/shirou/gopsutil/v3/host"
  "github.com/shirou/gopsutil/v3/load"
  "github.com/shirou/gopsutil/v3/process"
  "github.com/shirou/gopsutil/v3/net"

)


func init() {
  env.Packages["stdlib/system"] = map[string]reflect.Value{
    "Memory":             reflect.ValueOf(mem.VirtualMemory),
    "Cpu":                reflect.ValueOf(cpu.Info),
    "CpuUtil":            reflect.ValueOf(cpu.Times),
    "Partitions":         reflect.ValueOf(disk.Partitions),
    "Disk":               reflect.ValueOf(disk.Usage),
    "DiskIO":             reflect.ValueOf(disk.IOCounters),
    "Temperature":        reflect.ValueOf(host.SensorsTemperatures),
    "BootTime":           reflect.ValueOf(host.BootTime),
    "HostID":             reflect.ValueOf(host.HostID),
    "HostInfo":           reflect.ValueOf(host.Info),
    "Architecture":       reflect.ValueOf(host.KernelArch),
    "Kernel":             reflect.ValueOf(host.KernelVersion),
    "Platform":           reflect.ValueOf(host.PlatformInformation),
    "Uptime":             reflect.ValueOf(host.Uptime),
    "Virtualization":     reflect.ValueOf(host.Virtualization),
    "Load":               reflect.ValueOf(load.Avg),
    "Processes":          reflect.ValueOf(process.Processes),
    "Connections":        reflect.ValueOf(net.Connections),
    "ConntrackStats":     reflect.ValueOf(net.ConntrackStats),
    "Interfaces":         reflect.ValueOf(net.Interfaces),
  }
  env.PackageTypes["stdlib/system"] = map[string]reflect.Type{
    "Process":                  reflect.TypeOf(process.Process{}),
    "ConnectionStat":           reflect.TypeOf(net.ConnectionStat{}),
    "ConntrackStat":            reflect.TypeOf(net.ConntrackStat{}),
    "InterfaceStat":            reflect.TypeOf(net.InterfaceStat{}),
    "InfoStat":                 reflect.TypeOf(host.InfoStat{}),
  }
}
