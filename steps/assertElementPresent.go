package steps

import (
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["assertElementPresent"] = assertElementPresent
}

func assertElementPresent(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertElementPresent]: " + target)

	_, err := WD.FindElement(selenium.ByXPATH, target)

	if err != nil {
		AssertionFailure()
	}

	StepSuccess()
}
