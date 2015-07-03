package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertBodyText"] = assertBodyText
}

func assertBodyText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertBodyText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
