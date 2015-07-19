package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["assertPageSource"] = assertPageSource
}

func assertPageSource() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.documentElement.outerHTML;
	`

	fmt.Print("[assertPageSource]: " + Arg1)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	if strings.Contains(body, Arg2) {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
