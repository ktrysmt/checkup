package steps

import (
	"fmt"
)

func init() {
	StepList["get"] = get
}

func get(a interface{}) {

	url := a.(string)
	fmt.Print("[get]: " + url)
	err := WD.Get(url)
	StepFailure(err)

	StepSuccess()
}
