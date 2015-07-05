package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElement"] = verifyElement
}

func verifyElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyElement]: " + attr + " ")
	_, err := WD.FindElement(selenium.ByXPATH, attr)
	errors.VerifyResult(err)

}
