package steps

import (
	"fmt"
)

func init() {
	StepList["assertTitle"] = assertTitle
}

func assertTitle() {

	fmt.Print("[assertTitle]: " + Arg1)
	title, err := WD.Title()
	StepFailure(err)

	if title == Arg1 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
