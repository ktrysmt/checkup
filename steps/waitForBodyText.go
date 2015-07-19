package steps

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForBodyText"] = waitForBodyText
}

func waitForBodyText() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	fmt.Print("[waitForBodyText]: " + Arg1)

	arg := []interface{}{}
	limit := SetStepTimeout(Arg2)
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
