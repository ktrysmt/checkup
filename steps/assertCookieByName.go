package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertCookieByName"] = assertCookieByName
}

func assertCookieByName(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SymplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	fmt.Print("[assertCookieByName]: " + target + " => " + value)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == target && cookie.Value == value {
			StepSuccess()
			return
		}

	}

	AssertionFailure()
}
