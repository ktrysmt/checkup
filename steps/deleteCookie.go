package steps

import (
	"fmt"
)

func init() {
	StepList["deleteCookie"] = deleteCookie
}

func deleteCookie() {

	fmt.Print("[deleteCookie]: " + Arg1)

	err := WD.DeleteCookie(Arg1)
	StepFailure(err)

	StepSuccess()

}
