package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["print"] = print
}

func print(_url interface{}) {

	url := _url.(string)

	fmt.Print("[print]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
