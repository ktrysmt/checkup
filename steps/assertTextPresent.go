package steps

import (
	"fmt"
	"strings"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
)

func init() {
	StepList["assertTextPresent"] = assertTextPresent
}

func assertTextPresent(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[assertTextPresent]: " + target)

	arg := []interface{}{}
	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	m := strings.Index(body, target)

	if m != -1 {
		StepSuccess()
	} else {
		AssertionFailure()
	}

}
