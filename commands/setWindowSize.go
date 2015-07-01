package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["setWindowSize"] = setWindowSize
}

func setWindowSize(_url interface{}) {

	url := _url.(string)

	fmt.Print("[setWindowSize]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
