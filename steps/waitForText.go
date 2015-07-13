package steps

import (
	"checkup/errors"
	"fmt"
)

func init() {
	StepList["waitForText"] = waitForText
}

func waitForText(_url interface{}) {

	// todo
	url := _url.(string)

	fmt.Print("[waitForText]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
