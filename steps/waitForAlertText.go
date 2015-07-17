package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForAlertText"] = waitForAlertText
}

func waitForAlertText() {

	fmt.Print("[waitForAlertText]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		text, _ := WD.AlertText()
		if text == Arg1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
