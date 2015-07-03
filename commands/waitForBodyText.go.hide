package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForBodyText"] = waitForBodyText
}

func waitForBodyText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForBodyText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
