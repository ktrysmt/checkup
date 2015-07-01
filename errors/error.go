package errors

import (
	"fmt"
	"os"
)

func Fatal(err error) {

	if err != nil {
		//fmt.Println("Fatal Error: ", err)
		//os.Exit(0)
		//panic(err)
		fmt.Println(err)
		os.Exit(0)
	}

}

func Syntax(ok bool, message string) {
	if !ok {
		//panic(message)
		fmt.Println("Syntax Error:", message)
		os.Exit(0)
	}
}
