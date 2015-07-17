package steps

import (
	"fmt"
	"time"
)

func init() {
	StepList["waitForCookiePresent"] = waitForCookiePresent
}

func waitForCookiePresent() {

	fmt.Print("[waitForCookiePresent]: " + Arg1)

	limit := SetStepTimeout("")
	latency := 0
	for {
		if limit < latency {
			StepFailure("It reached a time-out.")
			break
		}

		c, _ := WD.GetCookies()

		for _, cookie := range c {
			if cookie.Name == Arg1 {
				StepSuccess()
				break
			}

		}

		time.Sleep(time.Millisecond * 500)
		latency += 500
	}
}
