package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["submitElement"] = submitElement
}

func submitElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[submitElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
