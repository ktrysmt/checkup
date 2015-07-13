package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElementSelected"] = verifyElementSelected
}

func verifyElementSelected(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyElementSelected]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	if ok == true {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
