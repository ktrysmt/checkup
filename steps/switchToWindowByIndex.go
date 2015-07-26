package steps

import (
	"fmt"
	"strconv"
)

func init() {
	StepList["switchToWindowByIndex"] = switchToWindowByIndex
}

func switchToWindowByIndex() {

	fmt.Print("[switchToWindowByIndex]: " + Arg1)

	number, err1 := strconv.Atoi(Arg1)
	StepFailure(err1)

	windows, err2 := WD.WindowHandles()
	StepFailure(err2)

	for index, _ := range windows {
		if index == number {
			err3 := WD.SwitchWindow(Arg1)
			StepFailure(err3)
			break
		}
	}

	StepSuccess()
}
