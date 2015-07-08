package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["setElementSelected"] = setElementSelected
}

func setElementSelected(a interface{}) {

	attr := a.(string)

	fmt.Print("[setElementSelected]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
