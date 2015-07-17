package steps

import (
	"fmt"
)

func init() {
	StepList["verifyEval"] = verifyEval
}

func verifyEval() {

	SCRIPT := SCRIPT_getElementsByXPath + Arg1

	fmt.Print("[verifyEval]: " + Arg1 + " => " + Arg2)

	arg := []interface{}{}
	res, err1 := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err1)

	r := SimplifyTypeAttributeValue(res)

	if r == Arg2 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
