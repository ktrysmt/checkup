package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["clickElement"] = clickElement
}

func clickElement(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[clickElement]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
