package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForAlertPresent"] = waitForAlertPresent
}

func waitForAlertPresent(a interface{}) {

	// set timeout
	limit := SetStepTimeout("")

	fmt.Print("[waitForAlertPresent]: ")

	// wait for
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		_, err := WD.AlertText()
		if err != nil {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}
}
