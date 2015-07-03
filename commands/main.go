package commands

import (
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

var WD selenium.WebDriver

var CommandList = map[string]func(interface{}){}

func Validate(datas map[interface{}]interface{}) {

	Dive("validate", datas)

}

func Do(remote string, datas map[interface{}]interface{}) {

	browser := datas[0].(map[interface{}]interface{})["testcase"].([]interface{})[0].(map[interface{}]interface{})["browser"].(string)

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
				if _, ok := CommandList[command]; !ok {
					errors.Syntax(ok, command+" is undefined command. ")
				}
			case "do":
				CommandList[command](args)
			}
		}
	}
}
