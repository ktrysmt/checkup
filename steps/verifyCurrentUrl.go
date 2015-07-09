package steps

import (
	"fmt"
)

func init() {
	StepList["verifyCurrentUrl"] = verifyCurrentUrl
}

func verifyCurrentUrl(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyCurrentUrl]: " + attr)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == attr {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
