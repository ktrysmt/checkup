package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeTitle"] = storeTitle
}

func storeTitle(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeTitle]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
