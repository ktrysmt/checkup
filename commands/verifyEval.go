package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyEval"] = verifyEval
}

func verifyEval(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyEval]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
