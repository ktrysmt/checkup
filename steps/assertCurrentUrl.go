package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertCurrentUrl"] = assertCurrentUrl
}

func assertCurrentUrl(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertCurrentUrl]: " + target)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == target {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
