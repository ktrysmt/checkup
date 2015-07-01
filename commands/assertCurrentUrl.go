package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertCurrentUrl"] = assertCurrentUrl
}

func assertCurrentUrl(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertCurrentUrl]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
