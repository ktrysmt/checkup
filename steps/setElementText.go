package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["setElementText"] = setElementText
}

func setElementText() {

	fmt.Print("[setElementText]: " + Arg1 + ", " + Arg2)

	elem, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	err2 := elem.Clear()
	StepFailure(err2)

	err3 := elem.SendKeys(Arg2)
	StepFailure(err3)

	StepSuccess()
}
