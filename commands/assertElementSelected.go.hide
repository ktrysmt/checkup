package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertElementSelected"] = assertElementSelected
}

func assertElementSelected(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertElementSelected]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
