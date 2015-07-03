package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForElementValue"] = waitForElementValue
}

func waitForElementValue(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForElementValue]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
