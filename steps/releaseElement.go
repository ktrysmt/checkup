package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["releaseElement"] = releaseElement
}

func releaseElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[releaseElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	err2 := btn.MoveTo(0, 0)
	StepFailure(err2)

	err3 := WD.ButtonUp()
	StepFailure(err3)

	StepSuccess()
}
