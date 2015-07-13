package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["verifyCurrentUrl"] = verifyCurrentUrl
}

func verifyCurrentUrl(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyCurrentUrl]: " + target)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == target {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
