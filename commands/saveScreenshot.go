package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"unidriver/errors"
)

func init() {
	CommandList["saveScreenshot"] = saveScreenshot
}

func saveScreenshot(_filename interface{}) {

	filename := _filename.(string)

	fmt.Print("[saveScreenshot]: " + filename + " ")
	binary, err := WD.Screenshot()
	errors.Fatal(err)
	ioutil.WriteFile(filename, binary, os.ModePerm)
	fmt.Print("SUCCESS")
	fmt.Println("")
}
