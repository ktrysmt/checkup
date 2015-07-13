package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["submitElement"] = submitElement
}

func submitElement(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[submitElement]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := btn.Submit()
	StepFailure(err2)

	StepSuccess()
}
