package steps

import (
	"fmt"
)

func init() {
	StepList["answerAlert"] = answerAlert
}

func answerAlert() {

	fmt.Print("[answerAlert]: " + Arg1)
	err := WD.SetAlertText(Arg1)
	StepFailure(err)

	StepSuccess()
}
