package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["dismissAlert"] = dismissAlert
}

func dismissAlert(_url interface{}) {

	url := _url.(string)

	fmt.Print("[dismissAlert]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
