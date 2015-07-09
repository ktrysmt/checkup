package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["assertPageSource"] = assertPageSource
}

func assertPageSource(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.documentElement.outerHTML;
	`

	attr := a.(string)

	fmt.Print("[assertPageSource]: " + attr)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	m := strings.Index(body, attr)

	if m != -1 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
