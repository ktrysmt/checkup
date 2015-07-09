package steps

import (
	"fmt"
)

func init() {
	StepList["assertCurrentUrl"] = assertCurrentUrl
}

func assertCurrentUrl(a interface{}) {

	attr := a.(string)

	fmt.Print("[assertCurrentUrl]: " + attr)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == attr {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
