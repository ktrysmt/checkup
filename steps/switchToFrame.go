package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["switchToFrame"] = switchToFrame
}

func switchToFrame() {

	fmt.Print("[switchToFrame]: " + Arg1)
	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	id, err2 := btn.GetAttribute("id")
	StepFailure(err2)

	err3 := WD.SwitchFrame(id)
	StepFailure(err3)

	StepSuccess()

}
