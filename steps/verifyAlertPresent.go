package steps

import (
	"fmt"
)

func init() {
	StepList["verifyAlertPresent"] = verifyAlertPresent
}

func verifyAlertPresent() {

	fmt.Print("[verifyAlertPresent]: ")
	_, err := WD.AlertText()

	if err != nil {
		VerificationFailure()
	} else {
		StepSuccess()
	}

}
