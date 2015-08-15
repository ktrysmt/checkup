package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"time"
)

func init() {
	StepList["waitForElementPresent"] = waitForElementPresent
}

func waitForElementPresent() {

	fmt.Print("[waitForElementPresent]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		_, err := WD.FindElement(SeleniumSelector, Arg1)
		if err != nil {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
