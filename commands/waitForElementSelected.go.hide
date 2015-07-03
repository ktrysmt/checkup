package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForElementSelected"] = waitForElementSelected
}

func waitForElementSelected(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForElementSelected]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
