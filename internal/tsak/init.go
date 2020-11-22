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
  "github.com/google/uuid"
  "github.com/erikdubbelboer/gspt"
  "github.com/common-nighthawk/go-figure"
)

func Init() {
  uid, _ := uuid.NewUUID()
  flag.BoolVar(&conf.Nocolor, "nocolor", false, "Disable colors in terminal output")
  flag.BoolVar(&conf.Debug, "debug", false, "Enable debug output")
  flag.BoolVar(&conf.Error, "error", false, "Enable ERROR outpout")
  flag.BoolVar(&conf.Warning, "warning", false, "Enable WARNING outpout")
  flag.BoolVar(&conf.Info, "info", false, "Enable INFO outpout")
  flag.BoolVar(&conf.Stdout, "stdout", false, "Send log entries to /dev/stdout as well")
  flag.BoolVar(&conf.TraceNR, "tracenr", false, "Send TSAK traces as logs to New Relic")
  flag.BoolVar(&conf.Production, "production", false, "Running Ushell in production mode")
  flag.StringVar(&conf.Nrapi, "nrapi", os.Getenv("NEW_RELIC_LICENSE_KEY"), "New Relic API key")
  flag.StringVar(&conf.Nrapiq, "nrapiq", os.Getenv("NEW_RELIC_Q_LICENSE_KEY"), "New Relic Query API key")
  flag.StringVar(&conf.Logfile, "log", "", "Name of the log file")
  flag.StringVar(&conf.ID, "id", uid.String(), "Application unique identifier")
  flag.StringVar(&conf.Name, "name", uid.String(), "Name of the application")
  flag.StringVar(&conf.Account, "account", os.Getenv("NEW_RELIC_ACCOUNT"), "New Relic user account number")
  flag.StringVar(&conf.Logapi, "logapi", "https://log-api.newrelic.com/log/v1", "LOG API endpoint")
  flag.StringVar(&conf.Evtapi, "evtapi", "https://insights-collector.newrelic.com/v1/accounts/%s/events", "EVT API endpoint")
  flag.StringVar(&conf.Queryapi, "queryapi", "https://insights-api.newrelic.com/v1/accounts/%s/query?nrql=%s", "Query API endpoint")
  flag.StringVar(&conf.Metricapi, "metricapi", "https://metric-api.newrelic.com/metric/v1", "Metric API endpoint")
  flag.IntVar(&conf.Maxsize, "logsize", 100, "Maximum size of the log file in Mb")
  flag.IntVar(&conf.Maxage, "logage", 7, "Maximum age of the logfile in days")
  flag.StringVar(&conf.In, "in", "", "Name of the script for the input")
  flag.StringVar(&conf.Proc, "proc", "", "Name of the script for the processing")
  flag.StringVar(&conf.Out, "out", "", "Name of the script for the output")
  flag.StringVar(&conf.Run, "run", "", "Name of the exclusive run script (-in/-out/-proc will be ignored)")
  flag.StringVar(&conf.Script, "script", "", "If passed, this script will be used as 'universal script' and will be passed to -in/-out/-proc/-housekeeper")
  flag.StringVar(&conf.Conf, "conf", "", "Configuration file")
  flag.StringVar(&conf.House, "housekeeper", "", "Housekeeper periodic script")
  flag.StringVar(&conf.Clips, "clips", "", "Name of non-exclusive main script executed in CLIPS environment")
  flag.Parse()
  if ! conf.Production {
    banner := figure.NewFigure("TSAK :> ", "", true)
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
  if conf.Script != "" {
    log.Trace(fmt.Sprintf("Universal script %s will be used for -in/-out/-proc/-housekeeper"))
  }
  if ! conf.Production {
    log.Trace(RunVerification())
  }
  signal.InitSignal()
  script.InitScript()
  clips.InitClips()
  log.Event("TsakEvent", logrus.Fields{
    "message":    "Application started",
    "evtc":       0,
  })
}
