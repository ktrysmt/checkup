package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyText"] = verifyText
}

func verifyText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
