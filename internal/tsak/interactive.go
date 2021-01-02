package tsak

import (
  "fmt"
  "github.com/common-nighthawk/go-figure"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/signal"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/script"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/conf"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/c-bata/go-prompt"
)

func tsak_completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "import", Description: "Importing TSAK modules"},
    {Text: "if", Description: "Do if statement"},
		{Text: "func", Description: "Define TSAK function"},
		{Text: "for", Description: "Do for loop"},
    {Text: "fmt.Println", Description: "Print a line"},
    {Text: "stdlib.ExitRequest()", Description: "Request an exit"},
    {Text: "true", Description: "Say truth"},
    {Text: "false", Description: "Do not say the truth"},
    {Text: "make", Description: "Make some structure or data item"},
    {Text: "struct", Description: "Define a structure"},
    {Text: "chan", Description: "Define a channel"},    
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func TsakShell() {
  log.Trace("Switching to interactive mode")
  banner := figure.NewFigure(fmt.Sprintf("TSAK %s:> ", conf.Ver), "", true)
  banner.Print()
  fmt.Println()
  conf.DisplayVersion()
  script.RunString("proc", `stdlib = import("stdlib")`)
  fmt.Println("run: stdlib.ExitRequest() for an exit.")
  for ! signal.ExitRequested() {
    code := prompt.Input(fmt.Sprintf("TSAK[%v]> ", conf.Name), tsak_completer)
    script.RunString("proc", code)
  }
  log.Trace("Exit out of interactive mode")
}
