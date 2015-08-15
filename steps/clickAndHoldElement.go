package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["clickAndHoldElement"] = clickAndHoldElement
}

func clickAndHoldElement() {

	fmt.Print("[clickAndHoldElement]: " + Arg1)

	btn, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.ButtonDown()
	StepFailure(err3)

	StepSuccess()
}
