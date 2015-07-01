package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertElementAttribute"] = assertElementAttribute
}

func assertElementAttribute(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertElementAttribute]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
