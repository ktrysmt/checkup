package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["verifyEval"] = verifyEval
}

func verifyEval(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	SCRIPT := SCRIPT_getElementsByXPath + target

	fmt.Print("[verifyEval]: " + target + " => " + value)

	arg := []interface{}{}
	res, err1 := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err1)

	r := SimplifyTypeAttributeValue(res)

	if r == value {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
