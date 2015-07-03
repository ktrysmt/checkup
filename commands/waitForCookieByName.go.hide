package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForCookieByName"] = waitForCookieByName
}

func waitForCookieByName(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForCookieByName]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
