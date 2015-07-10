package steps

import (
	"fmt"
)

func init() {
	StepList["goBack"] = goBack
}

func goBack(a interface{}) {

	fmt.Print("[goBack]: ")
	err := WD.Back()
	StepFailure(err)

	StepSuccess()

}
