package steps

import (
	"fmt"
)

func init() {
	StepList["assertEval"] = assertEval
}

func assertEval() {

	SCRIPT := SCRIPT_getElementsByXPath + Arg1

	fmt.Print("[assertEval]: " + Arg1 + ", " + Arg2)

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
