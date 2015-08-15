package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["clickElement"] = clickElement
}

func clickElement() {

	fmt.Print("[clickElement]: " + Arg1)

	btn, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
