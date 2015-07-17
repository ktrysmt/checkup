package steps

import (
	"fmt"
)

func init() {
	StepList["print"] = print
}

func print() {

	fmt.Print("[print]: " + Arg1)

}
