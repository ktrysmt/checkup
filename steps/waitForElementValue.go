package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"time"
)

func init() {
	StepList["waitForElementValue"] = waitForElementValue
}

func waitForElementValue() {

	fmt.Print("[waitForElementValue]: " + Arg1 + ", " + Arg2)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		elem, _ := WD.FindElement(selenium.ByXPATH, Arg1)

		text, _ := elem.GetAttribute("value")
		if text == Arg2 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
