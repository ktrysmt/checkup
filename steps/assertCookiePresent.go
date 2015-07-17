package steps

import (
	"fmt"
)

func init() {
	StepList["assertCookiePresent"] = assertCookiePresent
}

func assertCookiePresent() {

	fmt.Print("[assertCookiePresent]: " + Arg1)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == Arg1 {
			StepSuccess()
			return
		}

	}

	AssertionFailure()

}
