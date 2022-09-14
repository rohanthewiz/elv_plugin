package main

import (
	"github.com/rohanthewiz/rerr"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vars"
)

var Ns *eval.Ns

func init() {
	nb := eval.BuildNs()

	nb.AddVar("foo", vars.NewReadOnly("bar"))

	nb.AddGoFns(map[string]any{
		"hello": process,
	})

	Ns = nb.Ns()
}

func process(s string) string {
	str := rerr.FunctionLoc(1)
	return s + " " + str
}
