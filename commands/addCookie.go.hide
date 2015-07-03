package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["addCookie"] = addCookie
}

func addCookie(_url interface{}) {

	url := _url.(string)

	fmt.Print("[addCookie]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
