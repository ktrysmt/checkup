package steps

import (
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	StepList["saveScreenshot"] = saveScreenshot
}

func saveScreenshot() {

	fmt.Print("[saveScreenshot]: " + Arg1)
	binary, err := WD.Screenshot()

	if err == nil {
		err = ioutil.WriteFile(Arg1, binary, os.ModePerm)
		StepFailure(err)

	} else {
		StepFailure(err)
	}

	StepSuccess()
}
