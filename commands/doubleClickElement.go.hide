package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["doubleClickElement"] = doubleClickElement
}

func doubleClickElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[doubleClickElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
