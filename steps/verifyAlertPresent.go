package steps

import (
	"fmt"
)

func init() {
	StepList["verifyAlertPresent"] = verifyAlertPresent
}

func verifyAlertPresent(a interface{}) {

	//	attr := a.(string)

	fmt.Print("[verifyAlertPresent]: ")
	_, err := WD.AlertText()

	if err != nil {
		VerificationFailure()
	} else {
		StepSuccess()
	}

}
