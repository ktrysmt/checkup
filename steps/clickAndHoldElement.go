package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["clickAndHoldElement"] = clickAndHoldElement
}

func clickAndHoldElement(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[clickAndHoldElement]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.ButtonDown()
	StepFailure(err3)

	StepSuccess()
}
