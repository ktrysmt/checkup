package steps

import (
	"fmt"
)

func init() {
	StepList["assertAlertPresent"] = assertAlertPresent
}

func assertAlertPresent(a interface{}) {

	//attr := a.(string)

	fmt.Print("[assertAlertPresent] ")
	_, err := WD.AlertText()

	if err != nil {
		AssertionFailure()
	} else {
		StepSuccess()
	}

}
