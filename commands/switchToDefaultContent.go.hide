package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["switchToDefaultContent"] = switchToDefaultContent
}

func switchToDefaultContent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[switchToDefaultContent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
