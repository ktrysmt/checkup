package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertCookieByName"] = assertCookieByName
}

func assertCookieByName(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertCookieByName]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
