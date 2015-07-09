package steps

import (
	"fmt"
)

func init() {
	StepList["dismissAlert"] = dismissAlert
}

func dismissAlert(a interface{}) {

	fmt.Print("[dismissAlert]: ")
	err := WD.DismissAlert()
	StepFailure(err)

	StepSuccess()
}
