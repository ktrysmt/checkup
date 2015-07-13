package steps

import (
	"checkup/Godeps/_workspace/src/github.com/k0kubun/pp"
	"checkup/Godeps/_workspace/src/github.com/mattn/go-scan"
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"checkup/errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var WD selenium.WebDriver

var StepList = map[string]func(interface{}){}

var DefaultTimeout = "60000"

func Validate(datas map[interface{}]interface{}) {

	Dive("validate", datas)

}

func Do(remote string, datas map[interface{}]interface{}) {

	var browser string
	scan.ScanTree(datas, "/[0]/testcase[0]/browser/", &browser)

	caps := selenium.Capabilities{"browserName": browser}
	wd, err1 := selenium.NewRemote(caps, remote)
	WD = wd
	WDFatal(err1)
	defer WD.Quit()

	SetAroundTimeout(datas)
	/*
		var timeout time.Duration
		scan.ScanTree(datas, "/[0]/testcase[0]/timeout/", &timeout)
		if timeout == 0 {
			i, err2 := strconv.Atoi(DefaultTimeout)
			WDFatal(err2)
			timeout = time.Duration(i)
		} else {
			DefaultTimeout = strconv.Itoa(int(timeout))
		}
		WD.SetImplicitWaitTimeout(timeout)
	*/

	Dive("do", datas)

}

func Dive(flag string, datas map[interface{}]interface{}) {

	steps := datas[0].(map[interface{}]interface{})["testcase"].([]interface{})[0].(map[interface{}]interface{})["steps"].([]interface{})

	for _, ss := range steps {

		stepSet := NormalizeStepSet(ss)

		for s, a := range stepSet.(map[interface{}]interface{}) {

			step := SimplifyTypeAttributeTarget(s)
			args := NestingTarget(SimplifyTypeAttributeValue(a))

			switch flag {
			case "validate":
				_, ok := StepList[step]
				errors.Syntax(ok, "["+step+"] is undefined step.")

				_, t_ok := args.(map[interface{}]interface{})["target"]
				errors.Syntax(t_ok, "undefined the [target] attribute.")

				_, v_ok := args.(map[interface{}]interface{})["value"]
				if v_ok && !t_ok {
					errors.Syntax(false, "undefined [target] and [value] attributes.")
				}
			case "do":
				StepList[step](args)
			}
		}
	}
}

const SCRIPT_getElementsByXPath = `document.getElementsByXPath = function(expression, parentElement) {
  var r = [];
  var x = document.evaluate(expression, parentElement || document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);
  for (var i = 0, l = x.snapshotLength; i < l; i++) {
    r.push(x.snapshotItem(i));
  }
  return r;
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

//TODO:------------------------------------------------------
func PrintStep(name string, a ...interface{}) {

	var target, value interface{}

	ok1 := scan.ScanTree(a, "/[0]/target", &target)
	ok2 := scan.ScanTree(a, "/[0]/value", &value)
	pp.Println(a)
	pp.Println(target)
	pp.Println(ok1)
	pp.Println(value)
	pp.Println(ok2)
	os.Exit(9)

	for i, arg := range a {
		pp.Println("----")
		pp.Println(i)
		pp.Println(arg)
	}
	os.Exit(9)
}

func NestingTarget(o interface{}) interface{} {

	switch o.(type) {
	case int, string, bool:
		o = map[interface{}]interface{}{"target": o}
	}
	return o
}

func NormalizeStepSet(o interface{}) interface{} {

	var result map[interface{}]interface{}

	switch o.(type) {
	case map[interface{}]interface{}:
		result = o.(map[interface{}]interface{})
	default:
		result = map[interface{}]interface{}{o: nil}
	}

	return result

}

func SimplifyTypeAttributeTarget(o interface{}) string {
	var v string
	switch o.(type) {
	case float64:
		v = strconv.Itoa(int(o.(float64)))
	case float32:
		v = strconv.Itoa(int(o.(float32)))
	case int:
		v = strconv.Itoa(o.(int))
	case bool:
		v = strconv.FormatBool(o.(bool))
	case map[interface{}]interface{}:
		// none
	case nil:
		v = ""
	default:
		v = o.(string)
	}
	return v
}

func SimplifyTypeAttributeValue(o interface{}) interface{} {
	var v interface{}
	switch o.(type) {
	case float64:
		v = strconv.Itoa(int(o.(float64)))
	case float32:
		v = strconv.Itoa(int(o.(float32)))
	case int:
		v = strconv.Itoa(o.(int))
	case bool:
		v = strconv.FormatBool(o.(bool))
	case nil:
		v = ""
	default:
		v = o
	}
	return v
}

func StepFailure(err interface{}) {

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

func SetStepTimeout(value string) (int, string) {

	var view string

	if value == "" {
		value = DefaultTimeout
		view = value + " msec (default)"
	} else {
		view = value + " msec "
	}
	var err1 error
	limit, err1 := strconv.Atoi(value)
	StepFailure(err1)

	return limit, view
}

func SetAroundTimeout(datas map[interface{}]interface{}) {
	var timeout time.Duration
	scan.ScanTree(datas, "/[0]/testcase[0]/timeout/", &timeout)
	if timeout == 0 {
		i, err2 := strconv.Atoi(DefaultTimeout)
		WDFatal(err2)
		timeout = time.Duration(i)
	} else {
		DefaultTimeout = strconv.Itoa(int(timeout))
	}
	WD.SetImplicitWaitTimeout(timeout)

}
