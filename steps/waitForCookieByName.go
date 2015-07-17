package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForCookieByName"] = waitForCookieByName
}

func waitForCookieByName() {

	fmt.Print("[waitForCookieByName]: " + Arg1 + ", " + Arg2)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		c, _ := WD.GetCookies()
		for _, cookie := range c {
			if cookie.Name == Arg1 && cookie.Value == Arg2 {
				StepSuccess()
				return
			}

		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}

}
