package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["setElementText"] = setElementText
}

func setElementText(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SymplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	fmt.Print("[setElementText]: " + target + " => " + value)

	elem, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	err2 := elem.Clear()
	StepFailure(err2)

	err3 := elem.SendKeys(value)
	StepFailure(err3)

	StepSuccess()
}
