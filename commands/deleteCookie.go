package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["deleteCookie"] = deleteCookie
}

func deleteCookie(_url interface{}) {

	url := _url.(string)

	fmt.Print("[deleteCookie]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
