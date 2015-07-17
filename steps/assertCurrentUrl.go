package steps

import (
	"fmt"
)

func init() {
	StepList["assertCurrentUrl"] = assertCurrentUrl
}

func assertCurrentUrl() {

	fmt.Print("[assertCurrentUrl]: " + Arg1)
	url, err := WD.CurrentURL()
	StepFailure(err)

	if url == Arg1 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
