package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["store"] = store
}

func store(_url interface{}) {

	url := _url.(string)

	fmt.Print("[store]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
