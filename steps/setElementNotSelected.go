package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["setElementNotSelected"] = setElementNotSelected
}

func setElementNotSelected(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
var elem = document.getElementsByXPath(arguments[0]);
return elem[0].selected = false;`

	attr := a.(string)

	fmt.Print("[setElementNotSelected]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	var err3 error

	if ok {
		arg := []interface{}{attr}
		_, err3 = WD.ExecuteScript(SCRIPT, arg)
		StepFailure(err3)
	}

	StepSuccess()
}
