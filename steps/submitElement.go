package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["submitElement"] = submitElement
}

func submitElement() {

	fmt.Print("[submitElement]: " + Arg1)

	btn, err1 := WD.FindElement(selenium.ByXPATH, Arg1)
	StepFailure(err1)

	err2 := btn.Submit()
	StepFailure(err2)

	StepSuccess()
}
