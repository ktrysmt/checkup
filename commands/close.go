package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["close"] = close
}

func close(_url interface{}) {

	url := _url.(string)

	fmt.Print("[close]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
