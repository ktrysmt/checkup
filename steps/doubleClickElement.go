package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["doubleClickElement"] = doubleClickElement
}

func doubleClickElement(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[clickElement]: " + target + " ")

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.DoubleClick()
	StepFailure(err3)

	StepSuccess()

}
