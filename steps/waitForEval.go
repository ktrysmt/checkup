package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForEval"] = waitForEval
}

func waitForEval() {

	SCRIPT := SCRIPT_getElementsByXPath + Arg1

	fmt.Print("[waitForEval]: " + Arg1 + ", " + Arg2)

	arg := []interface{}{}
	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		r, _ := WD.ExecuteScript(SCRIPT, arg)
		response := SimplifyTypeAttributeValue(r)
		if response == Arg2 {
			StepSuccess()
			break
		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
