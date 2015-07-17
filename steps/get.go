package steps

import (
	"fmt"
	"regexp"
)

func init() {
	StepList["get"] = get
}

func get() {

	if m, _ := regexp.MatchString("^https?://", BaseUrl); m {
		Arg1 = BaseUrl + Arg1
	}

	fmt.Print("[get]: " + Arg1)
	err := WD.Get(Arg1)
	StepFailure(err)

	StepSuccess()
}
