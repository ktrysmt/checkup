package main

import (
	"io/ioutil"
	"os"
	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/Godeps/_workspace/src/gopkg.in/yaml"
	"unidriver/actions"
	"unidriver/errors"
)

type Parser struct {
	yamlfiles []string
}

func (p Parser) Do() map[interface{}]interface{} {

	datas := make(map[interface{}]interface{})
	data := make(map[interface{}]interface{})

	// parse yamls
	for index, yamlfile := range p.yamlfiles {

		// read yamlfile
		handle, err := os.Open(yamlfile)
		defer handle.Close()
		errors.Fatal(err)
		buffer, err := ioutil.ReadAll(handle)
		errors.Fatal(err)
		err = yaml.Unmarshal(buffer, &data)
		errors.Fatal(err)

		// is defined testsuites
		_, ok := data["testsuites"]
		errors.Syntax(ok, "undefined testsuites.")

		// is defined actions
		max := len(data["testsuites"].([]interface{}))
		for i := 0; i < max; i++ {
			_, ok := data["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
			errors.Syntax(ok, "undefined actions.")
		}

		datas[index] = data
	}

	// is actions
	actions.IsFunc()
	pp.Println("called IsFunc.")
	os.Exit(99999)
	/***
	for key, data := range datas {
		pp.Println(key)
		// pickup actions
		actions.isFunc()
		max := len(data["testsuites"].([]interface{}))
		//max := 1 // 1 yaml, 1 actions
		for i := 0; i < max; i++ {
			_, ok := data["testsuites"].([]interface{})[i].(map[interface{}]interface{})["actions"]
			errors.Syntax(ok, "undefined actions.")
		}

	}
	***/
	os.Exit(999)

	return datas
}
