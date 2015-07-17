package steps

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForPageSource"] = waitForPageSource
}

func waitForPageSource() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.documentElement.outerHTML;
	`

	fmt.Print("[waitForPageSource]: " + Arg1)

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
		if m := strings.Index(body, Arg1); m != -1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
