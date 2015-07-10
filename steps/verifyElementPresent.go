package steps

import (
	"fmt"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
)

func init() {
	StepList["verifyElementPresent"] = verifyElementPresent
}

func verifyElementPresent(a interface{}) {

	var target string
	scan.ScanTree(a, "/target", &target)

	fmt.Print("[verifyElementPresent]: " + target)

	_, err := WD.FindElement(selenium.ByXPATH, target)

	if err != nil {
		VerificationFailure()
	}

	StepSuccess()

}
