package steps

import (
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	StepList["saveScreenshot"] = saveScreenshot
}

func saveScreenshot(_filename interface{}) {

	filename := _filename.(string)

	fmt.Print("[saveScreenshot]: " + filename)
	binary, err := WD.Screenshot()

	if err == nil {
		err = ioutil.WriteFile(filename, binary, os.ModePerm)
		StepFailure(err)

	} else {
		StepFailure(err)
	}

	StepSuccess()
}
