package steps

import (
	"fmt"
)

func init() {
	StepList["verifyAlertPresent"] = verifyAlertPresent
}

func verifyAlertPresent(a interface{}) {

	fmt.Print("[verifyAlertPresent]: ")
	_, err := WD.AlertText()

	if err != nil {
		VerificationFailure()
	} else {
		StepSuccess()
	}

}
