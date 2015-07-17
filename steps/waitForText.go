package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForText"] = waitForText
}

func waitForText() {

	fmt.Print("[waitForText]: " + Arg1 + ", " + Arg2)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		elem, _ := WD.FindElement(selenium.ByXPATH, Arg1)

		text, _ := elem.Text()
		if m := strings.Index(text, Arg2); m != -1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
