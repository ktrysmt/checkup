package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertTitle"] = assertTitle
}

func assertTitle(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertTitle]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
