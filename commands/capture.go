package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"unidriver/errors"
)

func init() {
	CommandList["capture"] = capture
}

func capture(_filename interface{}) {

	filename := _filename.(string)

	fmt.Println("[capture]: " + filename)

	binary, err := WD.Screenshot()
	errors.Fatal(err)

	ioutil.WriteFile("./"+filename, binary, os.ModePerm)
}
