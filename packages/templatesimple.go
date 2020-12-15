package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  tpl "github.com/valyala/fasttemplate"
)

func init() {
  env.Packages["template/simple"] = map[string]reflect.Value{
    "Template":           reflect.ValueOf(tpl.ExecuteString),
  }
  env.PackageTypes["template/simple"] = map[string]reflect.Type{

  }
}
