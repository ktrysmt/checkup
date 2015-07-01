package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForAlertPresent"] = waitForAlertPresent
}

func waitForAlertPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForAlertPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
