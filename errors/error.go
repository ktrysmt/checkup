package errors

import (
	"fmt"
	"os"
)

func Fatal(err error) {

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

}

func Syntax(ok bool, message string) {

	if !ok {
		fmt.Println("Syntax Error:", message)
		os.Exit(0)
	}
}

func AssertResult(err error) {

	if err != nil {
		fmt.Print("FAILURE")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Print("SUCCESS")
		fmt.Println("")
	}
}

func VerifyResult(err error) {

	if err != nil {
		fmt.Print("FAILURE")
		fmt.Println("")
		fmt.Println(err)
	} else {
		fmt.Println("SUCCESS")
		fmt.Println("")
	}
}
