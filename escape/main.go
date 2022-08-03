package main

import (
	"bytes"
	"encoding/json"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	arg := wf.Args()[0]

	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(arg)
	if err != nil {
		panic(err)
	}
	v := buf.String()

	wf.NewItem(v).Subtitle("json escape").Arg(v).Valid(true)

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
