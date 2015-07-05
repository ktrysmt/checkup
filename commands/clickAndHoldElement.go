package commands

import (
	"fmt"
	//	"os"
	//	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["clickAndHoldElement"] = clickAndHoldElement
}

func clickAndHoldElement(a interface{}) {

	var attr string
	scan.ScanTree(a, "/target", &attr)

	fmt.Print("[clickAndHoldElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByCSSSelector, attr)
	errors.Fatal(err1)

	err2 := btn.MoveTo(0, 0)
	errors.Fatal(err2)

	err3 := WD.ButtonDown()
	errors.Fatal(err3)

	errors.PrintResult(err1, err2, err3)

}
