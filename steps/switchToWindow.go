package steps

import (
	"fmt"
)

func init() {
	StepList["switchToWindow"] = switchToWindow
}

func switchToWindow(a interface{}) {

	attr := a.(string)

	fmt.Print("[switchToWindow]: " + attr)
	err := WD.SwitchWindow(attr)
	StepFailure(err)

	StepSuccess()
}
