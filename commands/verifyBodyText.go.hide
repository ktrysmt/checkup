package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyBodyText"] = verifyBodyText
}

func verifyBodyText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyBodyText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
