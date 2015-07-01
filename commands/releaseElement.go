package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["releaseElement"] = releaseElement
}

func releaseElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[releaseElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
