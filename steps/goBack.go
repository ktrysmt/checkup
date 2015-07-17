package steps

import (
	"fmt"
)

func init() {
	StepList["goBack"] = goBack
}

func goBack() {

	fmt.Print("[goBack]: ")
	err := WD.Back()
	StepFailure(err)

	StepSuccess()

}
