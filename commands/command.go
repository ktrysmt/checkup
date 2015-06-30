package commands

import (
	//	"io/ioutil"
	//	"os"
	//	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

var WD selenium.WebDriver

var CommandList = map[string]func(interface{}){}

func Validate(datas map[interface{}]interface{}) {

	Dive("validate", datas)

}

func Do(browser string, remote string, datas map[interface{}]interface{}) {

	// create wd instance
	caps := selenium.Capabilities{"browserName": browser}
	wd, err := selenium.NewRemote(caps, remote)
	defer wd.Quit()
	WD = wd
	errors.Fatal(err)

	Dive("drive", datas)

}

func Dive(flag string, datas map[interface{}]interface{}) {

	// (*1) digest only 1 file yet.
	for _, data := range datas {
		// (*2) digest only 1 commands yet.
		max := len(data.(map[interface{}]interface{})["testsuites"].([]interface{}))
		for i := 0; i < max; i++ {
			commands, _ := data.(map[interface{}]interface{})["testsuites"].([]interface{})[i].(map[interface{}]interface{})["commands"].([]interface{})

			for _, commandSet := range commands {

				for _command, args := range commandSet.(map[interface{}]interface{}) {

					command := _command.(string)

					switch flag {
					case "validate":
						if _, ok := CommandList[command]; !ok {
							errors.Syntax(ok, command+" is undefined command. ")
						}
					case "drive":
						CommandList[command](args)
					}
				}
			}
			break // (*2)
		}
		break // (*1)
	}
}
