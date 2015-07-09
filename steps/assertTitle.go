package steps

import (
	"fmt"
)

func init() {
	StepList["assertTitle"] = assertTitle
}

func assertTitle(a interface{}) {

	attr := a.(string)

	fmt.Print("[assertTitle]: " + attr)
	title, err := WD.Title()
	StepFailure(err)

	if title == attr {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
