package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["setElementSelected"] = setElementSelected
}

func setElementSelected(a interface{}) {

	var attr string
	scan.ScanTree(a, "/target", &attr)

	fmt.Print("[setElementSelected]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	errors.Failure(err1)

	err2 := btn.Click()
	errors.Failure(err2)

	errors.Success(err1, err2)
}
