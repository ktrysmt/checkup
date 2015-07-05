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

	fmt.Print("[saveScreenshot]: " + filename)
	binary, err := WD.Screenshot()

	if err == nil {
		err = ioutil.WriteFile(filename, binary, os.ModePerm)
		errors.AssertResult(err)
	} else {
		errors.AssertResult(err)
	}
}
