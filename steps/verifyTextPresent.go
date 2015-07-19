package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["verifyTextPresent"] = verifyTextPresent
}

func verifyTextPresent() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	fmt.Print("[verifyTextPresent]: " + Arg1)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	if strings.Contains(body, Arg2) {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
