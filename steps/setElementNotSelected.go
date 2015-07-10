package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["setElementNotSelected"] = setElementNotSelected
}

func setElementNotSelected(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
var elem = document.getElementsByXPath(arguments[0]);
return elem[0].selected = false;`

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[setElementNotSelected]: " + target)

	btn, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	var err3 error

	if ok {
		arg := []interface{}{target}
		_, err3 = WD.ExecuteScript(SCRIPT, arg)
		StepFailure(err3)
	}

	StepSuccess()
}
