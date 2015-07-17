package steps

import (
	"fmt"
)

func init() {
	StepList["assertAlertText"] = assertAlertText
}

func assertAlertText() {

	fmt.Print("[assertAlertText]: " + Arg1)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == Arg1 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
