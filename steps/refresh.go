package steps

import (
	"fmt"
)

func init() {
	StepList["refresh"] = refresh
}

func refresh(a interface{}) {

	//attr := a.(string)

	fmt.Print("[refresh]: ")
	err := WD.Close()
	StepFailure(err)

	StepSuccess()

}
