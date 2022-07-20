package main

import (
	"strconv"
	"strings"

	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func lines(s string) string {
	n := strings.Count(s, "\n")
	if len(s) > 0 && !strings.HasSuffix(s, "\n") {
		n++
	}
	return strconv.Itoa(n)
}

func run() {
	arg := wf.Args()[0]

	wf.NewItem(strconv.Itoa(len(arg))).Subtitle("length").Arg(strconv.Itoa(len(arg))).Valid(true)
	wf.NewItem(lines(arg)).Subtitle("lines").Arg(lines(arg)).Valid(true)

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
