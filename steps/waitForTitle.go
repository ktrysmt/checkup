package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"fmt"
	"time"
)

func init() {
	StepList["waitForTitle"] = waitForTitle
}

func waitForTitle(a interface{}) {

	var t interface{}
	scan.ScanTree(a, "/target", &t)
	target := t.(string)

	// set timeout
	limit := SetStepTimeout("")

	fmt.Print("[waitForTitle]: " + target)

	// wait for
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		url, _ := WD.Title()
		if url == target {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
