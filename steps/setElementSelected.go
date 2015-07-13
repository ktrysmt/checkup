package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["setElementSelected"] = setElementSelected
}

func setElementSelected(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[setElementSelected]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
