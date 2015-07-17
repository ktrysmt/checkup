package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["assertElementPresent"] = assertElementPresent
}

func assertElementPresent() {

	fmt.Print("[assertElementPresent]: " + Arg1)

	_, err := WD.FindElement(selenium.ByXPATH, Arg1)

	if err != nil {
		AssertionFailure()
	}

	StepSuccess()
}
