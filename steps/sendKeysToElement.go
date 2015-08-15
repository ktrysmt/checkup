package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["sendKeysToElement"] = sendKeysToElement
}

func sendKeysToElement() {

	fmt.Print("[sendKeysToElement]: " + Arg1 + " => " + Arg2)

	elem, err1 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err1)

	err2 := elem.SendKeys(Arg2)
	StepFailure(err2)

	StepSuccess()

}
