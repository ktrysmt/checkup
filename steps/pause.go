package steps

import (
	"fmt"
	"strconv"
	"time"
)

func init() {
	StepList["pause"] = pause
}

func pause() {

	fmt.Print("[pause]: " + Arg1 + "(msec)")

	msec, err := strconv.Atoi(Arg1)
	StepFailure(err)

	time.Sleep(time.Millisecond * time.Duration(msec))

	StepSuccess()
}
