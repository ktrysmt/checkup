package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementPresent"] = verifyElementPresent
}

func verifyElementPresent(a interface{}) {

	attr := a.(string)

	fmt.Print("[verifyElementPresent]: " + attr + " ")

	_, err := WD.FindElement(selenium.ByXPATH, attr)
	errors.VerifyResult(err)

}
