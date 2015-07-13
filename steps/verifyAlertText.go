package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["verifyAlertText"] = verifyAlertText
}

func verifyAlertText(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyAlertText]: " + target)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == target {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
