package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["submitElement"] = submitElement
}

func submitElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[submitElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	err2 := btn.Submit()
	StepFailure(err2)

	StepSuccess()
}
