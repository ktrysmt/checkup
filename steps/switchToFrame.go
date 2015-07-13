package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["switchToFrame"] = switchToFrame
}

func switchToFrame(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[switchToFrame]: " + target)
	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	id, err2 := btn.GetAttribute("id")
	StepFailure(err2)

	err3 := WD.SwitchFrame(id)
	StepFailure(err3)

	StepSuccess()

}
