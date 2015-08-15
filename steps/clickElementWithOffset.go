package steps

import (
	//	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"fmt"
	"strconv"
)

func init() {
	StepList["clickElementWithOffset"] = clickElementWithOffset
}

func clickElementWithOffset() {

	fmt.Print("[clickElementWithOffset]: " + Arg1 + ", " + Arg2 + ", " + Arg3)

	w, err1 := strconv.Atoi(Arg2)
	StepFailure(err1)
	h, err2 := strconv.Atoi(Arg3)
	StepFailure(err2)

	elem, err3 := WD.FindElement(SeleniumSelector, Arg1)
	StepFailure(err3)

	err4 := elem.MoveTo(w, h)
	StepFailure(err4)

	err5 := elem.Click()
	StepFailure(err5)

	StepSuccess()
}
