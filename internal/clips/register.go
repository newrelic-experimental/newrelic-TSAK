package clips

import (
    "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
    "github.com/newrelic-experimental/newrelic-TSAK/internal/nr"
    "github.com/keysight/clipsgo/pkg/clips"
)

func RegisterVariables() {
  GetVarBindDef(clips.Symbol("INCH"), 0)
  GetVarBindDef(clips.Symbol("OUTCH"), 1)
  GetVarBindDef(clips.Symbol("CLIPS"), 2)
  GetVarBindDef(clips.Symbol("FACTS"), 3)
  GetVarBindDef(clips.Symbol("EVAL"), 4)
  GetVarBindDef(clips.Symbol("Answer"), 42)
}

func RegisterFunctions() {
  log.Trace("CLIPS functions registering")
  AddClipsFun("Now", nr.NowMillisec)
  AddClipsFun("exportallfacts", ExportAllFacts)
  AddClipsFun("exportassertedfacts", ExportAssertedFacts)
  AddClipsFun("var", SetVarBind)
  AddClipsFun("getvar", GetVarBind)
  AddClipsFun("VAR", GetVarBindDef)
  AddClipsFun("enablefactpipe", EnableFactPipe)
  AddClipsFun("disablefactpipe", DisableFactPipe)
  AddClipsFun("enablecmdpipe", EnableCmdPipe)
  AddClipsFun("disablecmdpipe", DisableCmdPipe)
  AddClipsFun("trace",    log.Trace)
  AddClipsFun("info",     log.Info)
  AddClipsFun("warning",  log.Warning)
  AddClipsFun("error",    log.Error)

  RegisterVariables()
}
