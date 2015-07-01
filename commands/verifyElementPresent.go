package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementPresent"] = verifyElementPresent
}

func verifyElementPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyElementPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
