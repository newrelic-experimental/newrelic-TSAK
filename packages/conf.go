package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
)

func ConfNRAPI(val string) string {
  conf.Nrapi = val
  return conf.Nrapi
}

func ConfNrapi() string {
  return conf.Nrapi
}

func ConfNRAPIQ(val string) string {
  conf.Nrapiq = val
  return conf.Nrapiq
}

func ConfNrapiq() string {
  return conf.Nrapiq
}

func ConfACCOUNT(val string) string {
  conf.Account = val
  return conf.Account
}

func ConfAccount() string {
  return conf.Account
}

func ConfArgs() []string {
  return conf.Args
}

func init() {
  env.Packages["conf"] = map[string]reflect.Value{
    "NRAPI":           reflect.ValueOf(ConfNRAPI),
    "Nrapi":           reflect.ValueOf(ConfNrapi),
    "NRAPIQ":          reflect.ValueOf(ConfNRAPIQ),
    "Nrapiq":          reflect.ValueOf(ConfNrapiq),
    "ACCOUNT":         reflect.ValueOf(ConfACCOUNT),
    "Account":         reflect.ValueOf(ConfAccount),
    "Args":            reflect.ValueOf(ConfArgs),
    "Version":         reflect.ValueOf(conf.Ver),
    "VersionMajor":    reflect.ValueOf(conf.VerMaj),
    "VersionMinor":    reflect.ValueOf(conf.VerMin),
    "VersionPrerelease":reflect.ValueOf(conf.VerPrerelease),
  }
  env.PackageTypes["conf"] = map[string]reflect.Type{

  }
}
