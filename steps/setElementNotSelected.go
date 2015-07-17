package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["setElementNotSelected"] = setElementNotSelected
}

func setElementNotSelected() {

	SCRIPT := SCRIPT_getElementsByXPath + `
var elem = document.getElementsByXPath(arguments[0]);
return elem[0].selected = false;`

	fmt.Print("[setElementNotSelected]: " + Arg1)

	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	ok, err2 := btn.IsSelected()
	StepFailure(err2)

	var err3 error

	if ok {
		arg := []interface{}{Arg1}
		_, err3 = WD.ExecuteScript(SCRIPT, arg)
		StepFailure(err3)
	}

	StepSuccess()
}
