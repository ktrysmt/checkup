package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"time"
)

func init() {
	StepList["waitForElementPresent"] = waitForElementPresent
}

func waitForElementPresent(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	// set timeout
	limit, view_value := SetStepTimeout(value)

	fmt.Print("[waitForElementPresent]: " + target + " : " + view_value)

	// wait for
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		_, err := WD.FindElement(selenium.ByXPATH, target)
		if err != nil {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
