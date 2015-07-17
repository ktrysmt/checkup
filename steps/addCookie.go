package steps

import (
	"checkup/Godeps/_workspace/src/github.com/tebeka/selenium"
	"encoding/json"
	"fmt"
)

func init() {
	StepList["addCookie"] = addCookie
}

func addCookie() {

	fmt.Print("[addCookie]: " + Arg1)

	var c selenium.Cookie
	var jsonString = []byte(Arg1)
	err1 := json.Unmarshal(jsonString, &c)
	StepFailure(err1)

	err2 := WD.AddCookie(&c)
	StepFailure(err2)

	StepSuccess()
}
