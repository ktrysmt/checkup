package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["assertElementSelected"] = assertElementSelected
}

func assertElementSelected() {

	fmt.Print("[assertElementSelected]: " + Arg1)

	btn, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	if ok == true {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
