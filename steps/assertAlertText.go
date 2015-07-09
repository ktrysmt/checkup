package steps

import (
	"fmt"
)

func init() {
	StepList["assertAlertText"] = assertAlertText
}

func assertAlertText(a interface{}) {

	attr := a.(string)

	fmt.Print("[assertAlertText]: " + attr)
	text, err := WD.AlertText()
	StepFailure(err)

	if text == attr {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
