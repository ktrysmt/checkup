package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["assertTextPresent"] = assertTextPresent
}

func assertTextPresent() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	fmt.Print("[assertTextPresent]: " + Arg1)

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
