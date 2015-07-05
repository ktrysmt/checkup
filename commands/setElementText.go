package commands

import (
	"fmt"
	//	"os"
	//	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["setElementText"] = setElementText
}

func setElementText(a interface{}) {

	var target, value string

	for key, str := range a.(map[interface{}]interface{}) {
		if key.(string) == "target" {
			target = str.(string)
		}
		if key.(string) == "value" {
			value = str.(string)
		}
	}

	fmt.Print("[setElementText]: " + target + " => " + value)
	elem, err := WD.FindElement(selenium.ByCSSSelector, target)

	if err == nil {
		elem.Clear()
		elem.SendKeys(value)
		errors.VerifyResult(err)
	} else {
		errors.AssertResult(err)
	}

}
