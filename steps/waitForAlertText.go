package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"fmt"
	"time"
)

func init() {
	StepList["waitForAlertText"] = waitForAlertText
}

func waitForAlertText(a interface{}) {

	var t interface{}
	scan.ScanTree(a, "/target", &t)
	target := t.(string)

	// set timeout
	limit := SetStepTimeout("")

	fmt.Print("[waitForAlertText]: " + target)

	// wait for
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		text, _ := WD.AlertText()
		if text == target {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
