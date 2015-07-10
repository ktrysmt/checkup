package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["get"] = get
}

func get(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[get]: " + target)
	err := WD.Get(target)
	StepFailure(err)

	StepSuccess()
}
