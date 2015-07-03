package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForText"] = waitForText
}

func waitForText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
