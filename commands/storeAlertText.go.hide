package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeAlertText"] = storeAlertText
}

func storeAlertText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeAlertText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
