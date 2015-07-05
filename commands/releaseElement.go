package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["releaseElement"] = releaseElement
}

func releaseElement(a interface{}) {

	var attr string
	scan.ScanTree(a, "/target", &attr)

	fmt.Print("[releaseElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByCSSSelector, attr)
	errors.Fatal(err1)

	err2 := btn.MoveTo(0, 0)
	errors.Fatal(err2)

	err3 := WD.ButtonUp()
	errors.Fatal(err3)

	errors.PrintResult(err1, err2, err3)
}
