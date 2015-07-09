package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["assertText"] = assertText
}

func assertText(a interface{}) {

	var target, value string

	scan.ScanTree(a, "/target", &target)
	scan.ScanTree(a, "/value", &value)

	fmt.Print("[assertText]: " + target + " => " + value)

	elem, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	text, err2 := elem.Text()
	StepFailure(err2)

	if text == value {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
