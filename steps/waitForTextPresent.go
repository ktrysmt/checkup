package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForTextPresent"] = waitForTextPresent
}

func waitForTextPresent(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	var t interface{}
	scan.ScanTree(a, "/target", &t)
	target := t.(string)

	// set timeout
	limit := SetStepTimeout("")

	fmt.Print("[waitForTextPresent]: " + target)

	// wait for
	arg := []interface{}{}
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		b, err3 := WD.ExecuteScript(SCRIPT, arg)
		StepFailure(err3)

		body := b.(string)
		if m := strings.Index(body, target); m != -1 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
