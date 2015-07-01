package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForAlertText"] = waitForAlertText
}

func waitForAlertText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForAlertText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
