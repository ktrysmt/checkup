package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertAlertPresent"] = assertAlertPresent
}

func assertAlertPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertAlertPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
