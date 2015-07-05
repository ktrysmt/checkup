package commands

import (
	//"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/mattn/go-scan"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

var WD selenium.WebDriver

var CommandList = map[string]func(interface{}){}

func Validate(datas map[interface{}]interface{}) {

	Dive("validate", datas)

}

func Do(remote string, datas map[interface{}]interface{}) {

	var browser string
	scan.ScanTree(datas, "/[0]/testcase[0]/browser/", &browser)

	caps := selenium.Capabilities{"browserName": browser}
	wd, err := selenium.NewRemote(caps, remote)
	WD = wd
	defer WD.Quit()
	errors.Fatal(err)

	Dive("do", datas)
}

func Dive(flag string, datas map[interface{}]interface{}) {

	commands := datas[0].(map[interface{}]interface{})["testcase"].([]interface{})[0].(map[interface{}]interface{})["commands"].([]interface{})

	for _, commandSet := range commands {

		for c, args := range commandSet.(map[interface{}]interface{}) {
			command := c.(string)

			switch flag {
			case "validate":
				_, ok := CommandList[command]
				errors.Syntax(ok, command+" is undefined command. ")

				// validate command args
				_, s := args.(string)
				_, i := args.(interface{})
				if s == false && i == true {
					_, target_ok := args.(map[interface{}]interface{})["target"]
					errors.Syntax(target_ok, "undefined the [target] attribute.")

					_, value_ok := args.(map[interface{}]interface{})["value"]
					if value_ok && !target_ok {
						errors.Syntax(false, "undefined [target] and [value] attributes.")
					}
				}
			case "do":
				CommandList[command](args)
			}
		}
	}
}
