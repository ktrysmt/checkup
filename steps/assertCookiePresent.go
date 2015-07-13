package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertCookiePresent"] = assertCookiePresent
}

func assertCookiePresent(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertCookiePresent]: " + target)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == target {
			StepSuccess()
			return
		}

	}

	AssertionFailure()

}
