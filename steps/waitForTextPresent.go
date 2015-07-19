package steps

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForTextPresent"] = waitForTextPresent
}

func waitForTextPresent() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	fmt.Print("[waitForTextPresent]: " + Arg1)

	arg := []interface{}{}
	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		b, _ := WD.ExecuteScript(SCRIPT, arg)
		body := b.(string)
		if strings.Contains(body, Arg1) {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
