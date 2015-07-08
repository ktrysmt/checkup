package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["doubleClickElement"] = doubleClickElement
}

func doubleClickElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[clickElement]: " + attr + " ")

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.DoubleClick()
	StepFailure(err3)

	StepSuccess()

}
