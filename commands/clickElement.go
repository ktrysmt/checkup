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

	fmt.Print("[clickElement]: " + attr)

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	errors.Fatal(err1)

	err2 := btn.Click()
	errors.Fatal(err2)

	errors.PrintResult(err1, err2)

}
