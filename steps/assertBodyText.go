package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["assertBodyText"] = assertBodyText
}

func assertBodyText() {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	fmt.Print("[assertBodyText]: " + Arg1)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	if strings.Contains(body, Arg1) {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
