package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForCookiePresent"] = waitForCookiePresent
}

func waitForCookiePresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForCookiePresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
