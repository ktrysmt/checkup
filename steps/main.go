package steps

import (
	"fmt"
	"os"
	"strconv"
	//	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

var WD selenium.WebDriver

var StepList = map[string]func(interface{}){}

func Validate(datas map[interface{}]interface{}) {
	Dive("validate", datas)

}

func Do(remote string, datas map[interface{}]interface{}) {

	var browser string
	scan.ScanTree(datas, "/[0]/testcase[0]/browser/", &browser)

	caps := selenium.Capabilities{"browserName": browser}
	wd, err := selenium.NewRemote(caps, remote)
	WD = wd
	WDFatal(err)

	Dive("do", datas)
}

func Dive(flag string, datas map[interface{}]interface{}) {

	steps := datas[0].(map[interface{}]interface{})["testcase"].([]interface{})[0].(map[interface{}]interface{})["steps"].([]interface{})

	for _, stepSet := range steps {

		for c, args := range stepSet.(map[interface{}]interface{}) {
			step := c.(string)
			_, n := args.(int)
			_, s := args.(string)
			_, i := args.(interface{})

			switch flag {
			case "validate":

				_, ok := StepList[step]
				errors.Syntax(ok, step+" is undefined step. ")

				if n == false && s == false && i == true {
					_, target_ok := args.(map[interface{}]interface{})["target"]
					errors.Syntax(target_ok, "undefined the [target] attribute.")

					_, value_ok := args.(map[interface{}]interface{})["value"]
					if value_ok && !target_ok {
						errors.Syntax(false, "undefined [target] and [value] attributes.")
					}
				}

			case "do":

				if n == false && s == false && i == true {
					target, _ := args.(map[interface{}]interface{})["target"]
					_, value_ok := args.(map[interface{}]interface{})["value"]
					if !value_ok {
						args = target
					}
				}
				if n == true && s == false && i == true {
					var err error
					args = strconv.Itoa(args.(int))
					WDFatal(err)
				}

				StepList[step](args)
			}
		}
	}
}

const SCRIPT_getElementsByXPath = `document.getElementsByXPath = function(expression, parentElement) {
  var r = []
  var x = document.evaluate(expression, parentElement || document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null)
  for (var i = 0, l = x.snapshotLength; i < l; i++) {
    r.push(x.snapshotItem(i))
  }
  return r
}
`

func WDFatal(err error) {
	if err != nil {
		fmt.Println("")
		fmt.Println(err)

		os.Exit(WDQuit())
	}
}

func WDQuit() int {
	defer WD.Quit()
	return 1
}

func StepFailure(err error) {

	if err != nil {
		fmt.Print(" Failure")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(WDQuit())
	}

}

func StepSuccess() {
	fmt.Print(" Success")
	fmt.Println("")
}

func AssertionFailure() {

	fmt.Print(" Assertion Failure")
	fmt.Println("")
	os.Exit(WDQuit())

}

func VerificationFailure() {

	fmt.Print(" Verification Failure")
	fmt.Println("")

}
