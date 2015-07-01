package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["clickAndHoldElement"] = clickAndHoldElement
}

func clickAndHoldElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[clickAndHoldElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
