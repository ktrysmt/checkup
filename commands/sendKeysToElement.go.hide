package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["sendKeysToElement"] = sendKeysToElement
}

func sendKeysToElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[sendKeysToElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
