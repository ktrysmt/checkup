package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertText"] = assertText
}

func assertText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
