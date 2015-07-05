package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["clickElement"] = clickElement
}

func clickElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[clickElement]: " + attr + " ")

	btn, err := WD.FindElement(selenium.ByCSSSelector, attr)

	if err == nil {
		err = btn.Click()
		errors.VerifyResult(err)
	} else {
		errors.AssertResult(err)
	}
}
