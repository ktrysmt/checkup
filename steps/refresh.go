package steps

import (
	"fmt"
)

func init() {
	StepList["refresh"] = refresh
}

func refresh(a interface{}) {

	fmt.Print("[refresh]: ")
	err := WD.Refresh()
	StepFailure(err)

	StepSuccess()

}
