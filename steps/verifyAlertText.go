package steps

import (
	"fmt"
)

func init() {
	StepList["verifyAlertText"] = verifyAlertText
}

func verifyAlertText(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyAlertText]: " + attr)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == attr {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
