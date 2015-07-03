package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyTextPresent"] = verifyTextPresent
}

func verifyTextPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyTextPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
