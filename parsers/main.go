package parsers

import (
	"checkup/Godeps/_workspace/src/gopkg.in/yaml"
	"checkup/errors"
	"fmt"
	"io/ioutil"
)

var (
	undefinedMessage = "Undefined [%s] in the file."
)

func ParseYaml(yamlfiles []string) map[interface{}]interface{} {

	datas := make(map[interface{}]interface{})

	// parse yamls
	for index, yamlfile := range yamlfiles {

		data, err := readYaml(yamlfile)
		errors.Fatal(err)

		ok, message := definedMetaDatas(data)
		errors.Syntax(ok, message)

		datas[index] = data
	}

	return datas
}

func readYaml(yamlfile string) (map[interface{}]interface{}, error) {

	data := make(map[interface{}]interface{})

	buf, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func definedMetaDatas(data map[interface{}]interface{}) (bool, string) {

	if _, ok := data["testcase"]; ok == false {
		return false, fmt.Sprintf(undefinedMessage, "testcase")
	}

	required := []string{"steps", "browser"}

	//max := len(data["testcase"].([]interface{}))
	//for i := 0; i < max; i++ {
	for _, name := range required {
		if _, ok := data["testcase"].([]interface{})[0].(map[interface{}]interface{})[name]; ok == false {
			return false, fmt.Sprintf(undefinedMessage, name)
		}
	}
	//}

	return true, ""
}
