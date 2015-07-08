package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["switchToFrame"] = switchToFrame
}

func switchToFrame(a interface{}) {

	attr := a.(string)

	fmt.Print("[switchToFrame]: " + attr)
	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	id, err2 := btn.GetAttribute("id")
	StepFailure(err2)

	err3 := WD.SwitchFrame(id)
	StepFailure(err3)

	StepSuccess()

}
