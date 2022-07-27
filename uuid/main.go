package main

import (
	"math/big"

	"github.com/deanishe/awgo"
	"github.com/google/uuid"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	u := uuid.New()

	var i big.Int
	i.SetBytes(u[:])
	b58 := i.Text(58)

	wf.NewItem(b58).Subtitle("base58").Arg(b58).Valid(true)
	wf.NewItem(u.String()).Subtitle("uuid4").Arg(u.String()).Valid(true)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
