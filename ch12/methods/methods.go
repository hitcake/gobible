package methods

import (
	"fmt"
	"reflect"
	"strings"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type: %v\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		methodName := t.Method(i).Name
		//fmt.Printf("method: %v, type: %v\n", methodName, methodType)
		fmt.Printf("func (%s) %s%s\n", t, methodName, strings.TrimPrefix(methodType.String(), "func"))
	}
}
