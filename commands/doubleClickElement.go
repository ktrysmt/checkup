package commands

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

func init() {
	CommandList["doubleClickElement"] = doubleClickElement
}

func doubleClickElement(a interface{}) {

	attr := a.(string)

	fmt.Print("[clickElement]: " + attr + " ")

	btn, err1 := WD.FindElement(selenium.ByXPATH, attr)
	errors.Fatal(err1)

	err2 := btn.MoveTo(0, 0)
	errors.Fatal(err2)

	err3 := WD.DoubleClick()
	errors.Fatal(err3)

	errors.PrintResult(err1, err2, err3)

}
