package steps

import (
	"fmt"
)

func init() {
	StepList["answerAlert"] = answerAlert
}

func answerAlert(a interface{}) {

	attr := a.(string)

	fmt.Print("[answerAlert]: " + attr)
	err := WD.SetAlertText(attr)
	StepFailure(err)

	StepSuccess()
}
