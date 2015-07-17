package steps

import (
	"fmt"
)

func init() {
	StepList["verifyCookiePresent"] = verifyCookiePresent
}

func verifyCookiePresent() {

	fmt.Print("[verifyCookiePresent]: " + Arg1)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == Arg1 {
			StepSuccess()
			return
		}

	}

	VerificationFailure()

}
