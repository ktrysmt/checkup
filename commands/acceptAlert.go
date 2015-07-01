package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["acceptAlert"] = acceptAlert
}

func acceptAlert(_url interface{}) {

	url := _url.(string)

	fmt.Print("[acceptAlert]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
