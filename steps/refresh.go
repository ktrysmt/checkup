package steps

import (
	"fmt"
)

func init() {
	StepList["refresh"] = refresh
}

func refresh() {

	fmt.Print("[refresh]: ")
	err := WD.Refresh()
	StepFailure(err)

	StepSuccess()

}
