package steps

import (
	"fmt"
)

func init() {
	StepList["verifyCookieByName"] = verifyCookieByName
}

func verifyCookieByName() {

	fmt.Print("[verifyCookieByName]: " + Arg1 + ", " + Arg2)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == Arg1 && cookie.Value == Arg2 {
			StepSuccess()
			return
		}

	}

	VerificationFailure()
}
