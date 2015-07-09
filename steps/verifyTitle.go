package steps

import (
	"fmt"
)

func init() {
	StepList["verifyTitle"] = verifyTitle
}

func verifyTitle(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyTitle]: " + attr)
	title, err := WD.Title()
	StepFailure(err)

	if title == attr {
		StepSuccess()
	} else {
		VerificationFailure()
	}
}
