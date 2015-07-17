package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["setElementSelected"] = setElementSelected
}

func setElementSelected() {

	fmt.Print("[setElementSelected]: " + Arg1)

	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
