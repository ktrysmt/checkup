package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["switchToFrameByIndex"] = switchToFrameByIndex
}

func switchToFrameByIndex(_url interface{}) {

	url := _url.(string)

	fmt.Print("[switchToFrameByIndex]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
