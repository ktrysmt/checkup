package steps

import (
	"fmt"
)

func init() {
	StepList["close"] = close
}

func close(a interface{}) {

	fmt.Print("[close]: ")
	err := WD.Close()
	StepFailure(err)

	StepSuccess()
}
