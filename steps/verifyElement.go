package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElement"] = verifyElement
}

func verifyElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyElement]: " + attr)
	_, err := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err)

	StepSuccess()

}
