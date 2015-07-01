package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["clickElement"] = clickElement
}

func clickElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[clickElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
