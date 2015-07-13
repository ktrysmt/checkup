package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["assertElementValue"] = assertElementValue
}

func assertElementValue(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	fmt.Print("[assertElementValue]: " + target + " => " + value)

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
