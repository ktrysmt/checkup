package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["doubleClickElement"] = doubleClickElement
}

func doubleClickElement() {

	fmt.Print("[clickElement]: " + Arg1)

	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.DoubleClick()
	StepFailure(err3)

	StepSuccess()

}
