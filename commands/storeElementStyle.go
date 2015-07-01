package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeElementStyle"] = storeElementStyle
}

func storeElementStyle(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeElementStyle]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
