package conf


type arrayFlags []string

func (i *arrayFlags) String() string {
    return ""
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

var Debug bool
var Error bool
var Info bool
var Warning bool
var Production bool
var Nocolor bool
var Stdout bool
var TraceNR bool
var TelemetryLog bool
var Logfile string
var Maxsize int
var Maxage int
var Timeout int
var Name string
var ID string
var Account string
var Nrapi string
var Nrapiq string
var Logapi string
var Evtapi string
var Metricapi string
var Queryapi string
var In string
var Proc string
var Out string
var Conf string
var Run string
var House string
var Clips string
var Script string
var Every int
var Hkeep int
var EventType string
var AppPath string
var TelemetryDB string
var IsVersion bool
var IsStop bool
var IsInteractive bool
var MetricsToNR bool
var IPv6 bool
// P2P options
var IsP2P bool
var P2PExternalAddress string
var P2PBind string
var P2PBindPort int
var P2PIdleTimeout int
var P2PMaxInbound int
var P2PMaxOutbound int
var P2PDiscovery int
var P2PBootstrap = make(arrayFlags, 0)

// Versions
var Ver = "0.5"
var VerMaj = 0
var VerMin = 5
var VerPrerelease = 0
var Args []string
