package log

import (
  "fmt"
  "os"
  "io"
  "github.com/sirupsen/logrus"
  "gopkg.in/natefinch/lumberjack.v2"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
)

type Fields logrus.Fields

var log = logrus.New()
var ljack *lumberjack.Logger
var writer io.Writer

func Event(msg string, ctx logrus.Fields) {
  Log().WithFields(ctx).Trace(msg)
  nr.Event(msg, ctx)
}

func Metric(mname string, _type string, value interface{}, ctx logrus.Fields) {
  Log().WithFields(ctx).Trace(fmt.Sprintf("Metric[%s](%s)=%v", mname, _type, value))
  nr.Metric(mname, _type, value, ctx)
}

func InitLog() {
  if len(conf.Logfile) > 0 {
    ljack = &lumberjack.Logger{
      Filename:   conf.Logfile,
      MaxSize:    conf.Maxsize,
      MaxAge:     conf.Maxage,
      Compress:   false,
    }

    if conf.Stdout {
        writer = io.MultiWriter(os.Stdout, ljack)
    } else {
       writer = io.MultiWriter(ljack)
    }
  } else {
    if conf.Stdout {
      writer = io.MultiWriter(os.Stdout)
    } else {
      writer = io.MultiWriter()
    }
  }
  if conf.Production {
    log.SetFormatter(&logrus.JSONFormatter{})
    log.Level = logrus.TraceLevel
    log.SetOutput(writer)
  } else {
    log.SetFormatter(&logrus.TextFormatter{})
    if conf.Nocolor {
      log.Formatter.(*logrus.TextFormatter).DisableColors = true
    } else {
      log.Formatter.(*logrus.TextFormatter).DisableColors = false
    }
    log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
    log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
    log.Level = logrus.TraceLevel
    log.SetOutput(writer)
  }
  if conf.Debug {
    conf.Info = true
    conf.Warning = true
    conf.Error = true
  }
  if conf.Info {
    conf.Info = true
    conf.Warning = true
    conf.Error = true
  }
  if conf.Warning {
    conf.Warning = true
    conf.Error = true
  }
  Trace(fmt.Sprintf("Production level: %v", conf.Production))
  Trace(fmt.Sprintf("Maximum size of log file (Mb): %v", conf.Maxsize))
  Trace(fmt.Sprintf("Maximum age of log file (days): %v", conf.Maxage))
  Trace(fmt.Sprintf("Application UUID %v", conf.ID))
  if conf.Nrapi != "" {
    Trace(fmt.Sprintf("NRAPI Enabled"))
  }
  Trace("Log subsystem initialized")
}


func Trace(msg string, ctx ...logrus.Fields) {
  var c logrus.Fields
  if conf.Debug {
    if len(ctx) > 0 {
      c = ctx[0]
    } else {
      c = logrus.Fields{}
    }
    Log().WithFields(c).Trace(msg)
    if conf.TraceNR {
      if conf.Nrapi != "" {
        if conf.Production {
          nr.Log(msg, "trace", c)
        }
      }
    }
  }
}

func Info(msg string, ctx ...logrus.Fields) {
  var c logrus.Fields
  if conf.Info {
    if len(ctx) > 0 {
      c = ctx[0]
    } else {
      c = logrus.Fields{}
    }
    Log().WithFields(c).Info(msg)
    if conf.Nrapi != "" {
      if conf.Production {
        nr.Log(msg, "info", c)
      }
    }
  }
}

func Warning(msg string, ctx ...logrus.Fields) {
  var c logrus.Fields
  if conf.Warning {
    if len(ctx) > 0 {
      c = ctx[0]
    } else {
      c = logrus.Fields{}
    }
    Log().WithFields(c).Warning(msg)
    if conf.Nrapi != "" {
      if conf.Production {
        nr.Log(msg, "warning", c)
      }
    }
  }
}

func Error(msg string, ctx ...logrus.Fields) {
  var c logrus.Fields
  if conf.Error {
    if len(ctx) > 0 {
      c = ctx[0]
    } else {
      c = logrus.Fields{}
    }
    Log().WithFields(c).Error(msg)
    if conf.Nrapi != "" {
      if conf.Production {
        nr.Log(msg, "error", c)
      }
    }
  }
}

func Log() *logrus.Logger {
    return log
}

func Shutdown() {
  Trace("Log subsystem shutodown")
}
