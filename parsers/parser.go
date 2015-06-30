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

		// is defined testsuites
		_, ok := data["testsuites"]
		errors.Syntax(ok, "undefined [testsuites] in the file.")

		// is defined commands
		max := len(data["testsuites"].([]interface{}))
		for i := 0; i < max; i++ {
			_, ok := data["testsuites"].([]interface{})[i].(map[interface{}]interface{})["commands"]
			errors.Syntax(ok, "undefined [commands] in the file.")
		}

		datas[index] = data
	}

	return datas
}
