package steps

import (
	"fmt"
	"strconv"
)

func init() {
	StepList["setWindowSize"] = setWindowSize
}

func setWindowSize() {

	fmt.Print("[setWindowSize]: " + Arg1 + ", " + Arg2 + ", " + Arg3)

	w, err1 := strconv.Atoi(Arg1)
	StepFailure(err1)
	h, err2 := strconv.Atoi(Arg2)
	StepFailure(err2)

	err3 := WD.ResizeWindow(Arg3, w, h)
	StepFailure(err3)

	StepSuccess()
}
