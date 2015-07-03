package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["goForward"] = goForward
}

func goForward(_url interface{}) {

	url := _url.(string)

	fmt.Print("[goForward]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
