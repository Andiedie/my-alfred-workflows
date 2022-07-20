package main

import (
	"encoding/json"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	arg := wf.Args()[0]

	v, _ := json.Marshal(arg)

	wf.NewItem(string(v)).Subtitle("json escape").Arg(string(v)).Valid(true)

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
