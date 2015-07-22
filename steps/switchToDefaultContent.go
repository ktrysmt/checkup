package steps

import (
	"fmt"
)

func init() {
	StepList["switchToDefaultContent"] = switchToDefaultContent
}

func switchToDefaultContent() {

	fmt.Print("[switchToDefaultContent]: " + Arg1)

	err := WD.SwitchFrame("")
	StepFailure(err)

	StepSuccess()

}
