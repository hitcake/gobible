package practice_2

import (
	"fmt"
	"gobible/ch12/format"
	"reflect"
)

var maxDepth = 5

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T)\n", name, x)
	depth := 1
	display(name, reflect.ValueOf(x), &depth, maxDepth)
}

func display(path string, v reflect.Value, depth *int, maxDepth int) {
	if *depth > maxDepth {
		fmt.Printf("%s = %s \n", path, format.Any(v.Interface()))
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s is invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			*depth++
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth, maxDepth)
			*depth--
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			*depth++
			display(fieldPath, v.Field(i), depth, maxDepth)
			*depth--
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			*depth++
			display(fmt.Sprintf("%s[%s]", path, format.Any(key.Interface())), v.MapIndex(key), depth, maxDepth)
			*depth--
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		} else {
			*depth++
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth, maxDepth)
			*depth--
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		} else {
			fmt.Printf("%s.type = %s \n", path, v.Elem().Type())
			*depth++
			display(path+".value", v.Elem(), depth, maxDepth)
			*depth--
		}
	default:
		fmt.Printf("%s = %s \n", path, format.Any(v.Interface()))

	}
}
