package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/slack-go/slack"
)

func init() {
  env.Packages["protocols/slack"] = map[string]reflect.Value{
    "New":                  reflect.ValueOf(slack.New),
    "MsgOptionText":        reflect.ValueOf(slack.MsgOptionText),
    "MsgOptionAttachments": reflect.ValueOf(slack.MsgOptionAttachments),
    "MsgOptionAsUser":      reflect.ValueOf(slack.MsgOptionAsUser),
  }
  env.PackageTypes["protocols/slack"] = map[string]reflect.Type{
    "Attachment":           reflect.TypeOf(slack.Attachment{}),
  }
}
