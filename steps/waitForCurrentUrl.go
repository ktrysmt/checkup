package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForCurrentUrl"] = waitForCurrentUrl
}

func waitForCurrentUrl() {

	fmt.Print("[waitForCurrentUrl]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		url, _ := WD.CurrentURL()
		if url == Arg1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
