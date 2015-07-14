package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"time"
)

func init() {
	StepList["waitForText"] = waitForText
}

func waitForText(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	// set timeout
	limit := SetStepTimeout("")

	fmt.Print("[waitForPageSource]: " + target + " => " + value)

	// wait for
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		elem, err1 := WD.FindElement(selenium.ByXPATH, target)
		StepFailure(err1)

		text, err2 := elem.Text()
		StepFailure(err2)

		if text == value {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
