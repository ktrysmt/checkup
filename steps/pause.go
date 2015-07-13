package steps

import (
	"fmt"
	"strconv"
	"time"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["pause"] = pause
}

func pause(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[pause]: " + target + "(msec)")

	msec, err := strconv.Atoi(target)
	StepFailure(err)

	time.Sleep(time.Millisecond * time.Duration(msec))

	StepSuccess()
}
