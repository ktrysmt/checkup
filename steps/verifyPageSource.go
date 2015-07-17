package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["verifyPageSource"] = verifyPageSource
}

func verifyPageSource() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.documentElement.outerHTML;
	`

	fmt.Print("[verifyPageSource]: " + Arg1)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	m := strings.Index(body, Arg1)

	if m != -1 {
		StepSuccess()
	} else {
		VerificationFailure()
	}
}
