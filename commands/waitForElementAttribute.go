package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForElementAttribute"] = waitForElementAttribute
}

func waitForElementAttribute(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForElementAttribute]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
