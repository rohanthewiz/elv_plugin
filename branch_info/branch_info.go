package branch_info

import (
	"bufio"
	"bytes"
	"log"
	"strings"

	"github.com/elves/sample-plugin/command"
	"github.com/rohanthewiz/rerr"
)

type BranchListItem struct {
	Name      string
	Details   string
	IsCurrent bool
}

// GetBranchList returns a list of local branches
func GetBranchList() (brList []BranchListItem, err error) {
	out, err := command.ExecCmd("git", "branch", "-vv")
	if err != nil {
		return brList, rerr.Wrap(err)
	}

	uniqBranches := make(map[string]struct{}, 16)

	scnr := bufio.NewScanner(bytes.NewReader(out))

	for scnr.Scan() { // each line
		isCurrent := false
		line := strings.TrimSpace(scnr.Text())

		bef, aft, fnd := strings.Cut(line, " ")
		if !fnd {
			log.Println("This one is weird:", line)
			continue
		}

		if strings.Contains(bef, "*") {
			bef, aft, fnd = strings.Cut(aft, " ")
			if !fnd {
				log.Println("No more spaces in this one:", aft)
				continue
			}
			isCurrent = true
		}

		br := bef

		if _, ok := uniqBranches[br]; ok { // if we already have it, continue
			continue
		}
		uniqBranches[br] = struct{}{} // track

		brList = append(brList, BranchListItem{Name: br, Details: aft, IsCurrent: isCurrent})
	}

	return brList, err
}
