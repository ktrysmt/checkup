package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeText"] = storeText
}

func storeText(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
