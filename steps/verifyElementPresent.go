package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["verifyElementPresent"] = verifyElementPresent
}

func verifyElementPresent() {

	fmt.Print("[verifyElementPresent]: " + Arg1)

	_, err := WD.FindElement(selenium.ByXPATH, Arg1)

	if err != nil {
		VerificationFailure()
	}

	StepSuccess()

}
