package steps

import (
	"fmt"
)

func init() {
	StepList["verifyTitle"] = verifyTitle
}

func verifyTitle() {

	fmt.Print("[verifyTitle]: " + Arg1)
	title, err := WD.Title()
	StepFailure(err)

	if title == Arg1 {
		StepSuccess()
	} else {
		VerificationFailure()
	}
}
