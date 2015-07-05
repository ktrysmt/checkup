package errors

import (
	"fmt"
	"os"
)

func Fatal(err error) {

	if err != nil {
		fmt.Println("")
		fmt.Println(err)
		//panic(err)
		os.Exit(0)
	}

}

func Syntax(ok bool, message string) {

	if !ok {
		fmt.Println("")
		fmt.Println("Syntax Error:", message)
		os.Exit(0)
	}
}

func AssertResult(err error) {

	if err != nil {
		fmt.Print(" FAILURE")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Print(" SUCCESS")
		fmt.Println("")
	}
}

func VerifyResult(err error) {

	if err != nil {
		fmt.Print(" FAILURE")
		fmt.Println("")
		fmt.Println(err)
	} else {
		fmt.Print(" SUCCESS")
		fmt.Println("")
	}
}

func PrintResult(err ...error) {

	for _, s := range err {
		if s != nil {
			fmt.Print(" FAILURE")
			fmt.Println("")
			return
		}
	}
	fmt.Print(" SUCCESS")
	fmt.Println("")
}

func Failure(err error) {

	if err != nil {
		fmt.Print(" FAILURE")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(0)
	}

}

func Success(err ...error) {
	fmt.Print(" SUCCESS")
	fmt.Println("")
}
