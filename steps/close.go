package steps

import (
	"fmt"
)

func init() {
	StepList["close"] = close
}

func close() {

	fmt.Print("[close]: ")
	err := WD.Close()
	StepFailure(err)

	StepSuccess()
}
