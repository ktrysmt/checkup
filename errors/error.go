package errors

import (
	"fmt"
	"os"
)

func Fatal(err error) {

	if err != nil {
		fmt.Println("")
		fmt.Fprintln(os.Stderr, err)
		//panic(err)
		os.Exit(1)
	}

}

func Syntax(ok bool, message string) {

	if !ok {
		fmt.Println("")
		fmt.Fprintln(os.Stderr, "Syntax Error: "+message)
		os.Exit(1)
	}
}
