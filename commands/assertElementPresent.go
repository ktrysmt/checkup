package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertElementPresent"] = assertElementPresent
}

func assertElementPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertElementPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
