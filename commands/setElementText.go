package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["setElementText"] = setElementText
}

func setElementText(a interface{}) {

	var target, value string

	scan.ScanTree(a, "/target", &target)
	scan.ScanTree(a, "/value", &value)

	fmt.Print("[setElementText]: " + target + " => " + value)

	elem, err1 := WD.FindElement(selenium.ByXPATH, target)
	errors.Fatal(err1)

	err2 := elem.Clear()
	errors.Fatal(err2)

	err3 := elem.SendKeys(value)
	errors.Fatal(err3)

	errors.PrintResult(err1, err2, err3)

}
