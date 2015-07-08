package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["assertElementPresent"] = assertElementPresent
}

func assertElementPresent(a interface{}) {

	attr := a.(string)

	fmt.Print("[assertElementPresent]: " + attr)

	_, err := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err)

	if err != nil {
		AssertionFailure()
	}

	StepSuccess()
}
