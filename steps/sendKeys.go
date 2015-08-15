package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
)

func init() {
	StepList["sendKeys"] = sendKeys
}

func sendKeys() {

	fmt.Print("[sendKeys]: " + Arg1)

	elem, err1 := WD.FindElement(SeleniumSelector, "/")
	StepFailure(err1)

	err2 := elem.SendKeys(Arg1)
	StepFailure(err2)

	StepSuccess()

}
