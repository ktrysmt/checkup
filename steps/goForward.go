package steps

import (
	"fmt"
)

func init() {
	StepList["goForward"] = goForward
}

func goForward() {

	fmt.Print("[goForward]: ")
	err := WD.Forward()
	StepFailure(err)

	StepSuccess()

}
