package main

import (
	"io/ioutil"
	"os"
	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/github.com/tebeka/selenium"
	"unidriver/errors"
)

type Driver struct {
	browser string
	remote  string
	datas   map[interface{}]interface{}
}

func (d Driver) Do() {

	// create wd instance
	caps := selenium.Capabilities{"browserName": d.browser}
	wd, err := selenium.NewRemote(caps, d.remote)
	defer wd.Quit()
	errors.Fatal(err)

	// (*1) digest only 1 file yet.
	for _, data := range d.datas {
		// (*2) digest only 1 actions yet.
		max := len(data.(map[interface{}]interface{})["testsuites"].([]interface{}))
		for i := 0; i < max; i++ {
			actions, _ := data.(map[interface{}]interface{})["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
			pp.Println(actions)

			//pp.Println(data)
			//pp.Println(actions)
			//actions, ok := data.([]interface{})["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
			//errors.Syntax(ok, "undefined actions.")

			//d := Driver{c.String("browser"), c.String("remote"), actions}
			//d.Do()

			break // (*2)
		}

		break // (*1)
	}

	os.Exit(0)

	// access each action
	/*******
	for _, val := range d.actions.([]interface{}) {
		//pp.Println(key)
		//pp.Println(val)
		for key := range val.(map[interface{}]interface{}) {
			pp.Println(key)
		}
		// visit
		if _, ok := val.(map[interface{}]interface{})["visit"]; ok {
			//	pp.Println(val.(map[interface{}]interface{})["visit"])
		}
	}
	******/

	os.Exit(0)

	// Get simple playground interface (google)
	wd.Get("https://google.co.jp/")
	binary, err := wd.Screenshot()
	errors.Fatal(err)

	ioutil.WriteFile("./test1.png", binary, os.ModePerm)

}
