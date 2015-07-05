package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["clearSelections"] = clearSelections
}

func clearSelections(a interface{}) {

	var attr string
	scan.ScanTree(a, "/target", &attr)

	fmt.Print("[clearSelections]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	errors.Fatal(err1)

	err2 := btn.Clear()
	errors.Fatal(err2)

	errors.PrintResult(err1, err2)
}
