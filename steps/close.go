package steps

import (
	"fmt"
)

func init() {
	StepList["close"] = close
}

func close(a interface{}) {

	//attr := a.(string)

	fmt.Print("[close]")
	err := WD.Close()
	StepFailure(err)

	StepSuccess()
}
