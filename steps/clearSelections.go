package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["clearSelections"] = clearSelections
}

func clearSelections() {

	SCRIPT := SCRIPT_getElementsByXPath + `
var elem = document.getElementsByXPath(arguments[0]);
var len = elem[0].length;
var i = 0;
for(i=0;len>i;i++){
  elem[0].options[i].selected = false;
}
return true;
`

	fmt.Print("[clearSelections]: " + Arg1)

	_, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	var err2 error

	arg := []interface{}{Arg1}
	_, err2 = WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err2)

	StepSuccess()
}
