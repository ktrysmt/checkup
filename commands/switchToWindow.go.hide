package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["switchToWindow"] = switchToWindow
}

func switchToWindow(_url interface{}) {

	url := _url.(string)

	fmt.Print("[switchToWindow]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
