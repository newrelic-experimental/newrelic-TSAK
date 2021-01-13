package snmp

import (
  "os"
  "path"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
  "encoding/json"
  dw "github.com/karrick/godirwalk"
  "github.com/hallidave/mibtool/smi"
  "github.com/Jeffail/gabs"
  "github.com/hjson/hjson-go"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
)

var mibs *smi.MIB

type SNMPConf struct {
  Mibpath string
  DBpath  string
  tbl     *gabs.Container
  sym     *gabs.Container
  ix      *gabs.Container
  discovery *gabs.Container
}

func InitMib(path string) {
  log.Trace(fmt.Sprintf("MIB's will be loaded from %s", path))
  mibs = smi.NewMIB(path)
}

func LoadModule(modname string) bool {
  if mibs == nil {
    log.Error("You have to initialize MIB")
    return false
  }
  err := mibs.LoadModules(modname)
  if err != nil {
    log.Error(fmt.Sprintf("Error loading %s: %s", modname, err))
    return false
  }
  log.Trace(fmt.Sprintf("Module loaded %s", modname))
  return true
}

func InitAndLoadAll(mibdirpath string) int {
  log.Trace(fmt.Sprintf("Loading and initializing MIB modules from: %s", mibdirpath))
  c := 0
  InitMib(mibdirpath)
  err := dw.Walk(mibdirpath, &dw.Options{
        Callback: func(osPathname string, de *dw.Dirent) error {
            fileStat, err := os.Stat(osPathname)
            if err != nil {
              return err
            }
            if fileStat.IsDir() {
              return nil
            }
            fn := path.Base(path.Clean(osPathname))
            fn = strings.TrimSuffix(fn, path.Ext(fn))
            log.Trace(fmt.Sprintf("Loading %s", fn))
            LoadModule(fn)
            c+=1
            return nil
        },
        Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
  })
  if err != nil {
    log.Error(fmt.Sprintf("Error scanning MIB tree: %s", err))
  }
  log.Trace(fmt.Sprintf("%d modules been loaded", c))
  return c
}

func LoadDBFile(tag string, dbpath string, filename string) (res *gabs.Container) {
  buf, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dbpath, filename))
  if err != nil {
    log.Error(fmt.Sprintf("Error reading %v DB: %v", tag, err))
    return nil
  }
  res, err = ParseDB(buf)
  if err != nil {
    log.Error(fmt.Sprintf("Error reading %v DB: %v", tag, err))
    return nil
  }
  return
}

func InitDB(mibdirpath string, dbpath string) *SNMPConf {
  var s SNMPConf
  s.Mibpath = mibdirpath
  s.DBpath = dbpath
  s.sym = LoadDBFile("SYMBOLS", dbpath, "symbols.table")
  s.tbl = LoadDBFile("TABLES", dbpath, "tables.table")
  s.ix = LoadDBFile("INDEX", dbpath, "ix.table")
  s.discovery = LoadDBFile("DISCOVERY", dbpath, "discovery.table")
  return &s
}

func (s *SNMPConf) Symbol(oid string) *gabs.Container {
  if s.sym != nil && s.sym.Exists(oid) {
    res := gabs.New()
    res.Set(oid, "OID")
    res.Set(s.sym.Search(oid, "name").Data(), "name")
    res.Set(s.sym.Search(oid, "mib").Data(), "mib")
    res.Set(s.sym.Search(oid, "device").Data(), "device")
    return res
  }
  return nil
}

func (s *SNMPConf) Resolve(what string) *gabs.Container {
  if s.sym != nil && s.ix != nil && s.tbl != nil {
    oid := OID(what)
    if oid != "" {
      res := s.Symbol(oid)
      res.Set(what, "symbol")
      return res
    }
    sym := SYMBOL(what)
    if s.ix.Exists(what) {
      oid = s.ix.Search(what).Data().(string)
      res := s.Symbol(oid)
      sym := SYMBOL(oid)
      if sym == "" {
        res.Set("UNKNOWN", "symbol")
      } else {
        res.Set(sym, "symbol")
      }
      return res
    }
    res := s.Symbol(what)
    if res == nil {
      res = gabs.New()
      res.Set(what, "name")
      res.Set(what, "OID")
      res.Set("UNKNOWN-MIB", "mib")
      res.Set("generic", "device")
      if sym == "" {
        res.Set("UNKNOWN", "symbol")
      } else {
        res.Set(sym, "symbol")
      }
      return res
    }
    if sym == "" {
      res.Set("UNKNOWN", "symbol")
    } else {
      res.Set(sym, "symbol")
    }
    return res
  }
  log.Error("SNMP configuration is not initialized")
  return nil
}

func InitSNMP(mibdirpath string, dbpath string) *SNMPConf {
  log.Trace(fmt.Sprintf("Loading DB files from %v", dbpath))
  InitAndLoadAll(mibdirpath)
  return InitDB(mibdirpath, dbpath)
}

func ParseDB(data []byte) (res *gabs.Container, err error) {
  var hdata map[string]interface{}

  res = nil
  hjson.Unmarshal([]byte(data), &hdata)
  b, err := json.Marshal(hdata)
  if err != nil {
    return
  }
  res, err = gabs.ParseJSON(b)
  return
}

func OID(sym string) string {
  if mibs == nil {
    log.Error("MIBS database is not initialized")
    return ""
  }
  oid, err := mibs.OID(sym)
  if err != nil {
    // log.Error(fmt.Sprintf("Error resolving OID %s: %s", sym, err))
    return ""
  }
  return oid.String()
}

func IsOID(sym string) string {
  if mibs == nil {
    log.Error("MIBS database is not initialized")
    return sym
  }
  oid, err := mibs.OID(sym)
  if err != nil {
    return sym
  }
  return oid.String()
}

func SYMBOL(oid string) string {
  if mibs == nil {
    log.Error("MIBS database is not initialized")
    return ""
  }
  tmp := strings.Split(oid, ".")
  values := make([]int, 0, len(tmp))
  for _, raw := range tmp {
    v, err := strconv.Atoi(raw)
    if err != nil {
        // log.Error(fmt.Sprintf("Error converting OID element %s: %s", oid, err))
        continue
    }
    values = append(values, v)
  }
  return mibs.SymbolString(values)
}
