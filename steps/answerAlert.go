package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["answerAlert"] = answerAlert
}

func answerAlert(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[answerAlert]: " + target)
	err := WD.SetAlertText(target)
	StepFailure(err)

	StepSuccess()
}
