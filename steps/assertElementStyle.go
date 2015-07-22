package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["assertElementStyle"] = assertElementStyle
}

func assertElementStyle() {

	fmt.Print("[assertElementStyle]: " + Arg1 + ", " + Arg2 + ", " + Arg3)

	elem, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	style, err2 := elem.CSSProperty(Arg2)
	StepFailure(err2)

	if style == Arg3 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
