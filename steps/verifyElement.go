package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["verifyElement"] = verifyElement
}

func verifyElement() {

	fmt.Print("[verifyElement]: " + Arg1)
	_, err := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err)

	StepSuccess()

}
