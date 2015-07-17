package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["dragToAndDropElement"] = dragToAndDropElement
}

func dragToAndDropElement() {

	fmt.Print("[dragToAndDropElement]: " + Arg1 + ", " + Arg2)

	btn_src, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	btn_dst, err2 := WD.FindElement(selenium.ByXPATH, Arg2)
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
