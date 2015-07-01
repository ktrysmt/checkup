package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["answerAlert"] = answerAlert
}

func answerAlert(_url interface{}) {

	url := _url.(string)

	fmt.Print("[answerAlert]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
