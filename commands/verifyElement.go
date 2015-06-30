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

	fmt.Println("[verifyElement]: " + attr)
	_, err := WD.FindElement(selenium.ByCSSSelector, attr)
	errors.Fatal(err)

}
