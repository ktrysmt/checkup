package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeAlertPresent"] = storeAlertPresent
}

func storeAlertPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeAlertPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
