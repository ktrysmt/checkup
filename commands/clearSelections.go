package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["clearSelections"] = clearSelections
}

func clearSelections(_url interface{}) {

	url := _url.(string)

	fmt.Print("[clearSelections]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
