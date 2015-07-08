package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["clickElement"] = clickElement
}

func clickElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[clickElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	StepFailure(err1)

	err2 := btn.Click()
	StepFailure(err2)

	StepSuccess()
}
