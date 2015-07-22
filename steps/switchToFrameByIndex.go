package steps

import (
	"fmt"
)

func init() {
	StepList["switchToFrameByIndex"] = switchToFrameByIndex
}

func switchToFrameByIndex() {

	fmt.Print("[switchToFrameByIndex]: " + Arg1)

	err2 := WD.SwitchFrame(Arg1)
	StepFailure(err2)

	StepSuccess()

}
