package steps

import (
	"fmt"
	"regexp"
)

func init() {
	StepList["verifyBodyText"] = verifyBodyText
}

func verifyBodyText(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	attr := a.(string)

	fmt.Print("[assertBodyText]: " + attr)

	arg := []interface{}{}

	b, err1 := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err1)

	body := b.(string)
	m, err2 := regexp.MatchString(attr, body)
	StepFailure(err2)

	if m {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
