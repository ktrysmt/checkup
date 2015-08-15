package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"time"
)

func init() {
	StepList["waitForElementSelected"] = waitForElementSelected
}

func waitForElementSelected() {

	fmt.Print("[waitForElementSelected]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		btn, _ := WD.FindElement(SeleniumSelector, Arg1)

		if ok, _ := btn.IsSelected(); ok {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
