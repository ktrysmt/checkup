package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["verifyCookiePresent"] = verifyCookiePresent
}

func verifyCookiePresent(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyCookiePresent]: " + target)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == target {
			StepSuccess()
			return
		}

	}

	VerificationFailure()

}
