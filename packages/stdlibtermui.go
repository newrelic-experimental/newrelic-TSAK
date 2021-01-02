package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)


func init() {
  env.Packages["stdlib/termui"] = map[string]reflect.Value{
    "New":              reflect.ValueOf(ui.Init),
    "Close":            reflect.ValueOf(ui.Close),
    "Render":           reflect.ValueOf(ui.Render),
    "PollEvents":       reflect.ValueOf(ui.PollEvents),
    "KeyboardEvent":    reflect.ValueOf(ui.KeyboardEvent),
    "NewParagraph":     reflect.ValueOf(widgets.NewParagraph),
    "NewBarChart":      reflect.ValueOf(widgets.NewBarChart),
    "NewPieChart":      reflect.ValueOf(widgets.NewPieChart),
    "ColorYellow":      reflect.ValueOf(ui.ColorYellow),
    "ColorWhite":       reflect.ValueOf(ui.ColorWhite),
    "ColorBlack":       reflect.ValueOf(ui.ColorBlack),
    "ColorGreen":       reflect.ValueOf(ui.ColorGreen),
    "ColorBlue":        reflect.ValueOf(ui.ColorBlue),
    "ColorCyan":        reflect.ValueOf(ui.ColorCyan),
    "ColorMagenta":     reflect.ValueOf(ui.ColorMagenta),
  }
  env.PackageTypes["stdlib/termui"] = map[string]reflect.Type{
    "PieChart":         reflect.TypeOf(widgets.PieChart{}),
    "Paragraph":        reflect.TypeOf(widgets.Paragraph{}),
    "BarChart":         reflect.TypeOf(widgets.BarChart{}),
  }
}
