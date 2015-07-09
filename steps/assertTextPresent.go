package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["assertTextPresent"] = assertTextPresent
}

func assertTextPresent(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	attr := a.(string)

	fmt.Print("[assertTextPresent]: " + attr)

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
