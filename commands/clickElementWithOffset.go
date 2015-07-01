package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["clickElementWithOffset"] = clickElementWithOffset
}

func clickElementWithOffset(_url interface{}) {

	url := _url.(string)

	fmt.Print("[clickElementWithOffset]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
