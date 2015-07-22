package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["mouseOverElement"] = mouseOverElement
}

func mouseOverElement() {

	fmt.Print("[mouseOverElement]: " + Arg1)

	elem, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	err2 := elem.MoveTo(0, 0)
	StepFailure(err2)

	StepSuccess()
}
