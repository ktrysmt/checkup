package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["clearSelections"] = clearSelections
}

func clearSelections(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
var elem = document.getElementsByXPath(arguments[0]);
var len = elem[0].length;
var i = 0;
for(i=0;len>i;i++){
  elem[0].options[i].selected = false;
}
return true;
`
	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[clearSelections]: " + target)

	_, err1 := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err1)

	var err2 error

	arg := []interface{}{target}
	_, err2 = WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err2)

	StepSuccess()
}
