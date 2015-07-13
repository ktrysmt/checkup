package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"fmt"
	"strings"
	"time"
)

func init() {
	StepList["waitForPageSource"] = waitForPageSource
}

func waitForPageSource(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.documentElement.outerHTML;
	`

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	// set timeout
	limit, view_value := SetStepTimeout(value)

	fmt.Print("[waitForPageSource]: " + target + " : " + view_value)

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
