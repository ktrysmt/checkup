package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElementPresent"] = verifyElementPresent
}

func verifyElementPresent(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyElementPresent]: " + attr)

	_, err := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err)

	StepSuccess()

}
