package steps

import (
	"fmt"
)

func init() {
	StepList["verifyAlertText"] = verifyAlertText
}

func verifyAlertText() {

	fmt.Print("[verifyAlertText]: " + Arg1)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == Arg2 {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
