package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElementSelected"] = verifyElementSelected
}

func verifyElementSelected(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyElementSelected]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	if ok == true {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
