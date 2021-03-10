package packages

import (
  "reflect"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
  "github.com/google/gopacket/ip4defrag"
  "github.com/google/gopacket/layers"
  "github.com/mattn/anko/env"
)

func PcapNew(iface string, filter string) (*gopacket.PacketSource, error) {
  r, err := pcap.OpenLive(iface, 65536, true, pcap.BlockForever)
  if err != nil {
    return nil, err
  }
  err = r.SetBPFFilter(filter)
  if err != nil {
    return nil, err
  }
  return gopacket.NewPacketSource(r, r.LinkType()), nil
}


func init() {
  env.Packages["protocols/pcap"] = map[string]reflect.Value{
    "New":                    reflect.ValueOf(PcapNew),
    "Defragmenter":           reflect.ValueOf(ip4defrag.NewIPv4Defragmenter),
    "LayerIPv4":              reflect.ValueOf(layers.LayerTypeIPv4),
    "LayerTCP":               reflect.ValueOf(layers.LayerTypeTCP),
    "IPv4":                   reflect.ValueOf(layers.IPProtocolIPv4),
    "TCP":                    reflect.ValueOf(layers.IPProtocolTCP),
  }
  env.PackageTypes["protocols/pcap"] = map[string]reflect.Type{
    "Handle":                 reflect.TypeOf(pcap.Handle{}),
    "CaptureInfo":            reflect.TypeOf(gopacket.CaptureInfo{}),
    "PacketSource":           reflect.TypeOf(gopacket.PacketSource{}),
    "IPv4Defragmenter":       reflect.TypeOf(ip4defrag.IPv4Defragmenter{}),
    "Layer_IPv4":             reflect.TypeOf(layers.IPv4{}),
    "Flow":                   reflect.TypeOf(gopacket.Flow{}),
  }
}
