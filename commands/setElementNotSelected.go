package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["setElementNotSelected"] = setElementNotSelected
}

func setElementNotSelected(a interface{}) {

	var attr string
	scan.ScanTree(a, "/target", &attr)

	fmt.Print("[setElementNotSelected]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	errors.Fatal(err1)

	var err2, err3, err4, err5 error
	var ok bool

	ok, err2 = btn.IsSelected()
	errors.Fatal(err2)

	if ok {
		err3 = WD.SendModifier("ControlKey", true)
		errors.Fatal(err3)
		err4 = btn.Click()
		errors.Fatal(err4)
		err5 = WD.SendModifier("ControlKey", false)
		errors.Fatal(err5)
	}

	errors.PrintResult(err1, err2, err3, err4, err5)
}
