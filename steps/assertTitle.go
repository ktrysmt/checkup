package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertTitle"] = assertTitle
}

func assertTitle(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertTitle]: " + target)
	title, err := WD.Title()
	StepFailure(err)

	if title == target {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
