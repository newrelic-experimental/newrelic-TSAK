package tsak

import (
  "fmt"
  "flag"
  "os"
  "github.com/sirupsen/logrus"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/script"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/clips"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/telemetrydb"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/piping"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/snmpdb"
  "github.com/google/uuid"
  "github.com/erikdubbelboer/gspt"
  "github.com/common-nighthawk/go-figure"
)

func Init() {
  uid, _ := uuid.NewUUID()
  flag.BoolVar(&conf.Nocolor, "nocolor", false, "Disable colors in terminal output")
  flag.BoolVar(&conf.IsVersion, "version", false, "Display TSAK version")
  flag.BoolVar(&conf.IsInteractive, "interactive", false, "Run TSAK in interactive mode")
  flag.BoolVar(&conf.IPv6, "ipv6", false, "Enable IPv6 support")
  flag.StringVar(&conf.AppPath, "appdir", "/tmp", "Directory for an application and temporary files")
  flag.StringVar(&conf.TelemetryDB, "telemetrydb", ":memory:", "DSN string for opening an embedded telemetry DB")
  flag.BoolVar(&conf.Debug, "debug", false, "Enable debug output")
  flag.BoolVar(&conf.Error, "error", false, "Enable ERROR outpout")
  flag.BoolVar(&conf.Warning, "warning", false, "Enable WARNING outpout")
  flag.BoolVar(&conf.Info, "info", false, "Enable INFO outpout")
  flag.BoolVar(&conf.Stdout, "stdout", false, "Send log entries to /dev/stdout as well")
  flag.BoolVar(&conf.TraceNR, "tracenr", false, "Send TSAK traces as logs to New Relic")
  flag.BoolVar(&conf.Production, "production", false, "Running Ushell in production mode")
  flag.BoolVar(&conf.IsStop, "stop", false, "If -stop is passed, tsak will try to kill tsak process with same ID")
  flag.BoolVar(&conf.MetricsToNR, "nrspy", false, "Send an internal TSAK performance and stability metrics to New Relic")
  flag.BoolVar(&conf.TelemetryLog, "notelemetrylog", true, "Disable sending TSAK logs to telemetrydb")
  flag.StringVar(&conf.Nrapi, "nrapi", os.Getenv("NEW_RELIC_LICENSE_KEY"), "New Relic Insert API key")
  flag.StringVar(&conf.Nrapiq, "nrapiq", os.Getenv("NEW_RELIC_Q_LICENSE_KEY"), "New Relic Query API key")
  flag.StringVar(&conf.Logfile, "log", "", "Name of the log file")
  flag.StringVar(&conf.ID, "id", uid.String(), "Application unique identifier")
  flag.StringVar(&conf.Name, "name", uid.String(), "Name of the application")
  flag.StringVar(&conf.Account, "account", os.Getenv("NEW_RELIC_ACCOUNT"), "New Relic user account number")
  flag.StringVar(&conf.Logapi, "logapi", "https://log-api.newrelic.com/log/v1", "LOG API endpoint")
  flag.StringVar(&conf.Evtapi, "evtapi", "https://insights-collector.newrelic.com/v1/accounts/%s/events", "EVT API endpoint")
  flag.StringVar(&conf.Queryapi, "queryapi", "https://insights-api.newrelic.com/v1/accounts/%s/query?nrql=%s", "Query API endpoint")
  flag.StringVar(&conf.Metricapi, "metricapi", "https://metric-api.newrelic.com/metric/v1", "Metric API endpoint")
  flag.StringVar(&conf.EventType, "evttype", "TsakEvent", "Event type for the events generated by this instance")
  flag.IntVar(&conf.Maxsize, "logsize", 100, "Maximum size of the log file in Mb")
  flag.IntVar(&conf.Timeout, "timeout", 3, "Number of seconds for a network operatipons timeout")
  flag.IntVar(&conf.Hkeep, "hkeep", 0, "Minimum number of recent records to keep in History telemetrydb. 0 - unlimited")
  flag.IntVar(&conf.Maxage, "logage", 7, "Maximum age of the logfile in days")
  flag.IntVar(&conf.Every, "every", 200, "Number of items in consumption reporting batch")
  flag.StringVar(&conf.In, "in", "", "Name of the script for the input")
  flag.StringVar(&conf.Proc, "proc", "", "Name of the script for the processing")
  flag.StringVar(&conf.Out, "out", "", "Name of the script for the output")
  flag.StringVar(&conf.Run, "run", "", "Name of the exclusive run script (-in/-out/-proc will be ignored)")
  flag.StringVar(&conf.Script, "script", "", "If passed, this script will be used as 'universal script' and will be passed to -in/-out/-proc/-housekeeper")
  flag.StringVar(&conf.Conf, "conf", "", "Configuration file")
  flag.StringVar(&conf.House, "housekeeper", "", "Housekeeper periodic script")
  flag.StringVar(&conf.Clips, "clips", "", "Name of non-exclusive main script executed in CLIPS environment")
  // P2P flags
  flag.StringVar(&conf.P2PExternalAddress, "p2pexternal", "", "External address for incoming P2P traffic")
  flag.StringVar(&conf.P2PBind, "p2pbind", "127.0.0.1", "Bind address for P2P traffic")
  flag.IntVar(&conf.P2PBindPort, "p2pport", 10090, "Bind port for P2P traffic")
  flag.IntVar(&conf.P2PIdleTimeout, "p2ptimeout", 3, "P2P idle timeout")
  flag.IntVar(&conf.P2PMaxInbound, "p2pmaxinbound", 10000, "Maximum number of P2P inbound connections")
  flag.IntVar(&conf.P2PMaxOutbound, "p2pmaxoutbound", 10000, "Maximum number of P2P outbound connections")
  flag.IntVar(&conf.P2PDiscovery, "p2pdiscovery", 30, "Seconds between P2P discovery")
  flag.BoolVar(&conf.IsP2P, "p2p", false, "Enable P2P capabilities")
  flag.Var(&conf.P2PBootstrap, "bootstrap",  "P2P bootstrap nodes")
  // End of flags
  flag.Parse()
  if conf.In == "" && conf.Out == "" && conf.Proc == "" && conf.Script == "" && conf.Run == "" && ! conf.IsStop && ! conf.IsVersion && ! conf.IsInteractive {
    banner := figure.NewFigure(fmt.Sprintf("TSAK %s:> ", conf.Ver), "", true)
    banner.Print()
    fmt.Println()
    flag.PrintDefaults()
    fmt.Println("ERROR: You did not specified any of the TSAK command-line parameters...")
    os.Exit(0)
  }
  if conf.IsVersion {
    banner := figure.NewFigure(fmt.Sprintf("TSAK %s:> ", conf.Ver), "", true)
    banner.Print()
    fmt.Println()
    conf.DisplayVersion()
    fmt.Println("A bit of TSAK wisdom: ", RunVerification())
    os.Exit(0)
  }
  if ! conf.Production {
    banner := figure.NewFigure(fmt.Sprintf("TSAK %s:> ", conf.Ver), "", true)
    banner.Print()
    fmt.Println()
  }
  gspt.SetProcTitle(fmt.Sprintf("TSAK: %s[%s]", conf.Name, conf.ID))
  if conf.Script != "" {
    conf.In     = conf.Script
    conf.Out    = conf.Script
    conf.Proc   = conf.Script
    conf.House  = conf.Script
  }
  log.InitLog()
  log.Info(fmt.Sprintf("Application ID %v", conf.ID))
  if conf.IsStop {
    StopWithPid()
    os.Exit(0)
  } else {
    StartWithPid()
  }
  if conf.Script != "" {
    log.Trace(fmt.Sprintf("Universal script %s will be used for -in/-out/-proc/-housekeeper", conf.Script))
  }
  if ! conf.Production {
    log.Trace(RunVerification())
  }
  signal.InitSignal()
  piping.Init()
  script.InitScript()
  clips.InitClips()
  telemetrydb.Telemetrydb_Init()
  snmpdb.Init()
  InitP2P()
  log.Trace(fmt.Sprintf("TelemetryDB opens at: %v", conf.TelemetryDB))
  log.Trace(fmt.Sprintf("SQLITE version is %v", telemetrydb.TDBVersion()))
  if flag.NArg() > 0 {
    log.Trace(fmt.Sprintf("%v positional arguments have been passed to TSAK", flag.NArg()))
    for _, a := range flag.Args() {
      conf.Args = append(conf.Args, a)
    }
  }
  log.Info(fmt.Sprintf("Events generated by this TSAK instance will be of type: %v", conf.EventType))
  log.Event("TsakEvent", logrus.Fields{
    "message":    "Application started",
    "evtc":       0,
  })
}
