package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForCurrentUrl"] = waitForCurrentUrl
}

func waitForCurrentUrl(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForCurrentUrl]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
