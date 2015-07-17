package steps

import (
	"fmt"
)

func init() {
	StepList["switchToWindow"] = switchToWindow
}

func switchToWindow() {

	fmt.Print("[switchToWindow]: " + Arg1)
	err := WD.SwitchWindow(Arg1)
	StepFailure(err)

	StepSuccess()
}
