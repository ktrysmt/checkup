package parsers

import (
	"io/ioutil"
	"os"
	"unidriver/Godeps/_workspace/src/gopkg.in/yaml"
	"unidriver/errors"
)

func ParseYaml(yamlfiles []string) map[interface{}]interface{} {

	datas := make(map[interface{}]interface{})
	data := make(map[interface{}]interface{})

	// parse yamls
	for index, yamlfile := range yamlfiles {

		// read yamlfile
		handle, err := os.Open(yamlfile)
		defer handle.Close()
		errors.Fatal(err)
		buffer, err := ioutil.ReadAll(handle)
		errors.Fatal(err)
		err = yaml.Unmarshal(buffer, &data)
		errors.Fatal(err)

		// is defined testcase
		_, ok := data["testcase"]
		errors.Syntax(ok, "undefined [testcase] in the file.")

		// is defined steps
		max := len(data["testcase"].([]interface{}))
		for i := 0; i < max; i++ {
			var ok bool
			_, ok = data["testcase"].([]interface{})[i].(map[interface{}]interface{})["steps"]
			errors.Syntax(ok, "undefined [steps] in the file.")
			_, ok = data["testcase"].([]interface{})[i].(map[interface{}]interface{})["browser"]
			errors.Syntax(ok, "undefined [browser] in the file.")
		}

		datas[index] = data
	}

	return datas
}
