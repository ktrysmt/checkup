package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementStyle"] = verifyElementStyle
}

func verifyElementStyle(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyElementStyle]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
