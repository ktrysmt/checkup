package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["releaseElement"] = releaseElement
}

func releaseElement() {

	fmt.Print("[releaseElement]: " + Arg1)

	btn, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.ButtonUp()
	StepFailure(err3)

	StepSuccess()
}
