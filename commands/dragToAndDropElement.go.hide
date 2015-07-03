package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["dragToAndDropElement"] = dragToAndDropElement
}

func dragToAndDropElement(_url interface{}) {

	url := _url.(string)

	fmt.Print("[dragToAndDropElement]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
