package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["verifyCookieByName"] = verifyCookieByName
}

func verifyCookieByName(a interface{}) {

	var t, v interface{}

	scan.ScanTree(a, "/target", &t)
	scan.ScanTree(a, "/value", &v)
	val := SimplifyTypeAttributeValue(v)

	target := t.(string)
	value := val.(string)

	fmt.Print("[verifyCookieByName]: " + target + " => " + value)

	c, err := WD.GetCookies()
	StepFailure(err)

	for _, cookie := range c {
		if cookie.Name == target && cookie.Value == value {
			StepSuccess()
			return
		}

	}

	VerificationFailure()
}
