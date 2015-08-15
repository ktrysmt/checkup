package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["assertText"] = assertText
}

func assertText() {

	fmt.Print("[assertText]: " + Arg1 + " => " + Arg2)

	elem, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	text, err2 := elem.Text()
	StepFailure(err2)

	if text == Arg2 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
