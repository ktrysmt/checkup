package main

import (
	"checkup/Godeps/_workspace/src/github.com/codegangsta/cli"
	"checkup/parsers"
	"checkup/steps"
	"fmt"
	"os"
	"regexp"
)

func main() {

	cli.AppHelpTemplate = AppHelpTemplate
	app := cli.NewApp()
	app.Name = "checkup"
	app.Version = Version
	app.HideHelp = true
	app.Usage = "E2E Testing Application"
	app.Author = "aqafiam"
	app.Before = doBefore
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		remoteFlag,
	}
	app.Run(os.Args)

}

func doBefore(c *cli.Context) error {

	args := c.Args()

	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}
	for _, arg := range args {
		ok, _ := regexp.MatchString(`\.ya?ml$`, arg)
		if !ok {
			fmt.Println(arg + " is not yaml file.")
			os.Exit(2)
		}
	}

	return nil
}

func doMain(c *cli.Context) {

	steps.Datas = parsers.ParseYaml(c.Args())
	for _, data := range steps.Datas {
		steps.Init(c.String("remote"), data)
	}

}

var remoteFlag = cli.StringFlag{
	Name:  "remote, r",
	Value: "http://localhost:4444/wd/hub",
	Usage: "Remote WebDriver Server's URL",
}

const Version string = "0.1.0"

const AppHelpTemplate string = `Usage: {{.Name}} [options] [testcase.yml...]

Options:
{{range .Flags}}  {{.}}
{{end}}
`
