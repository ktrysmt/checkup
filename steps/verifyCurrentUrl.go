package steps

import (
	"fmt"
)

func init() {
	StepList["verifyCurrentUrl"] = verifyCurrentUrl
}

func verifyCurrentUrl() {

	fmt.Print("[verifyCurrentUrl]: " + Arg1)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == Arg1 {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
