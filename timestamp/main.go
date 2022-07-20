package main

import (
	"strconv"
	"time"

	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	arg := wf.Args()[0]
	var t time.Time
	if arg == "" {
		t = time.Now()
	} else if ts, err := strconv.ParseInt(arg, 10, 64); err == nil {
		var nanoSecond int64
		switch len(arg) {
		case 10:
			nanoSecond = ts * int64(time.Second)
		case 13:
			nanoSecond = ts * int64(time.Millisecond)
		case 16:
			nanoSecond = ts * int64(time.Microsecond)
		case 19:
			nanoSecond = ts * int64(time.Nanosecond)
		}
		t = time.Unix(0, nanoSecond)
	} else if tt, err := time.Parse("2006-01-02 15:04:05", arg); err == nil {
		t = tt
	}

	if !t.IsZero() {
		wf.NewItem(t.Format("2006-01-02 15:04:05")).Subtitle("formatted").Arg(t.Format("2006-01-02 15:04:05")).Valid(true)
		wf.NewItem(strconv.FormatInt(t.Unix(), 10)).Subtitle("second").Arg(strconv.FormatInt(t.Unix(), 10)).Valid(true)
		wf.NewItem(strconv.FormatInt(t.UnixMilli(), 10)).Subtitle("millisecond").Arg(strconv.FormatInt(t.UnixMilli(), 10)).Valid(true)
		wf.NewItem(strconv.FormatInt(t.UnixMicro(), 10)).Subtitle("microsecond").Arg(strconv.FormatInt(t.UnixMicro(), 10)).Valid(true)
		wf.NewItem(strconv.FormatInt(t.UnixNano(), 10)).Subtitle("nanosecond").Arg(strconv.FormatInt(t.UnixNano(), 10)).Valid(true)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
