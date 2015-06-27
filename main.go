package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	"regexp"
	"unidriver/Godeps/_workspace/src/github.com/codegangsta/cli"
	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	//	"unidriver/Godeps/_workspace/src/gopkg.in/yaml"
	//	"unidriver/errors"
)

func main() {
	app := cli.NewApp()
	app.Name = "unidriver"
	app.Version = "0.1.0"
	app.Usage = ""
	app.Author = "aqafiam"
	app.Before = doBefore
	app.Action = doMain
	app.Flags = []cli.Flag{
		browserFlag,
		remoteFlag,
	}
	app.Run(os.Args)
}

func doBefore(c *cli.Context) error {
	return nil
}

func doMain(c *cli.Context) {

	// show help
	args := c.Args()
	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	// check file name
	var yamlfiles []string
	for _, arg := range args {
		ok, _ := regexp.MatchString(".ya?ml$", arg)
		if !ok {
			fmt.Println(arg + "is not yaml file.")
			cli.ShowAppHelp(c)
			os.Exit(0)
		} else {
			yamlfiles = append(yamlfiles, arg)
		}
	}

	// open and read yamlfiles
	p := Parser{yamlfiles}
	datas := p.Do()

	// drive on datas
	d := Driver{c.String("browser"), c.String("remote"), datas}
	d.Do()

	os.Exit(999)

	// digest only 1 file yet.
	for _, data := range datas {
		// digest only 1 actions yet.
		for i := 0; i < 1; i++ {
			pp.Println(i)
			//pp.Println(data)
			_, ok := data.(map[interface{}]interface{})["testsuites"]
			//pp.Println(actions)
			pp.Println(ok)
			//actions, ok := data.([]interface{})["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
			//errors.Syntax(ok, "undefined actions.")

			//d := Driver{c.String("browser"), c.String("remote"), actions}
			//d.Do()
		}
	}

	// pickup actions
	//max := len(data["testsuites"].([]interface{}))
	//for i := 0; i < max; i++ {
	//for i := 0; i < 1; i++ {
	//	actions, ok := data["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
	//	errors.Syntax(ok, "undefined actions.")

	//	d := Driver{c.String("browser"), c.String("remote"), actions}
	//	d.Do()
	//}

}

var browserFlag = cli.StringFlag{
	Name:  "browser, b",
	Value: "firefox",
	Usage: "browser name",
}

var remoteFlag = cli.StringFlag{
	Name:  "remote, r",
	Value: "http://localhost:4444/wd/hub",
	Usage: "RemoteWebDriver server URL",
}
