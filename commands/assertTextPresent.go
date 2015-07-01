package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertTextPresent"] = assertTextPresent
}

func assertTextPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertTextPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
