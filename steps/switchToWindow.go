package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["switchToWindow"] = switchToWindow
}

func switchToWindow(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[switchToWindow]: " + target)
	err := WD.SwitchWindow(target)
	StepFailure(err)

	StepSuccess()
}
