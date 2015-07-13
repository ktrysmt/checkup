package steps

import (
	"fmt"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElement"] = verifyElement
}

func verifyElement(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyElement]: " + target)
	_, err := WD.FindElement(selenium.ByXPATH, target)
	StepFailure(err)

	StepSuccess()

}
