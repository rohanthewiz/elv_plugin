package main

import (
	"fmt"
	"strings"

	"github.com/elves/sample-plugin/branch_info"
	"github.com/rohanthewiz/rerr"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vars"
)

// Ns our plugin namespace
var Ns *eval.Ns

func init() {
	nb := eval.BuildNs()

	nb.AddVar("foo", vars.NewReadOnly("bar"))

	nb.AddGoFns(map[string]any{
		"hello": process,
	})

	Ns = nb.Ns()
}

func process() (ret string) {
	brs, err := branch_info.GetBranchList()
	if err != nil {
		fmt.Printf("error: %s", rerr.StringFromErr(err))
		return
	}

	fmt.Println(strings.Repeat("-", 70))

	for i, br := range brs {
		fmt.Printf("%d. %s %s - %s\n", i+1,
			func() (s string) {
				if br.IsCurrent {
					return "[current]"
				}
				return
			}(),
			br.Name, br.Details)
		fmt.Println(strings.Repeat("-", 70))
	}

	/* Hmm, how to stop and interactively pickup a value?
	var a any
	_, _ = fmt.Scanln(&a)
	*/
	return
}
