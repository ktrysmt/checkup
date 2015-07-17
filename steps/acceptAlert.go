package steps

import (
	"fmt"
)

func init() {
	StepList["acceptAlert"] = acceptAlert
}

func acceptAlert() {

	fmt.Print("[acceptAlert]: ")
	err := WD.AcceptAlert()
	StepFailure(err)

	StepSuccess()

}
