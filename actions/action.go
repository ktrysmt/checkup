package actions

import (
	"fmt"
	"os"
	//	"reflect"
	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
)

type Functions struct {
}

func (f *Functions) visit() {
	fmt.Println("calling visit. and I will visit the website.")
}

func IsFunc() {

	var f Functions

	fmt.Println("calling IsFunc.")

	//	functions := map[string]func(){
	//		"visit": visit,
	//	}

	a_act := "visit"
	pp.Println(f)
	pp.Println(a_act)
	//os.Exit(999999)
	//pp.Println(reflect.ValueOf(&f).MethodByName(a_act))
	//os.Exit(88)
	//reflect.ValueOf(&f).MethodByName(a_act).Call([]reflect.Value{})

	os.Exit(88)

}
