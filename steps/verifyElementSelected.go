package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["verifyElementSelected"] = verifyElementSelected
}

func verifyElementSelected() {

	fmt.Print("[verifyElementSelected]: " + Arg1)

	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	if ok == true {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
