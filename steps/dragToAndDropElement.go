package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["dragToAndDropElement"] = dragToAndDropElement
}

func dragToAndDropElement(a interface{}) {

	var target, value string
	scan.ScanTree(a, "/target", &target)
	scan.ScanTree(a, "/value", &value)

	fmt.Print("[dragToAndDropElement]: " + target + " => " + value)

	btn_src, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	btn_dst, err2 := WD.FindElement(selenium.ByXPATH, value)
	StepFailure(err2)

	err3 := btn_src.MoveTo(0, 0)
	StepFailure(err3)

	err4 := WD.ButtonDown()
	StepFailure(err4)

	err5 := btn_dst.MoveTo(0, 0)
	StepFailure(err5)

	err6 := WD.ButtonUp()
	StepFailure(err6)

	StepSuccess()
}
