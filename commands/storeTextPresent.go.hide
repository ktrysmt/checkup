package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeTextPresent"] = storeTextPresent
}

func storeTextPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeTextPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
