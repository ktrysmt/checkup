package steps

import (
	"fmt"
	"strconv"
	"time"
)

func init() {
	StepList["pause"] = pause
}

func pause(a interface{}) {

	target := a.(string)

	fmt.Print("[pause]: " + target + " msec")

	msec, err := strconv.Atoi(target)
	StepFailure(err)

	time.Sleep(time.Millisecond * time.Duration(msec))

	StepSuccess()
}
