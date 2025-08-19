package main

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

/*
编写相应的Pack函数，给定一个结构体值，Pack函数将返回合并了所有结构体成员和值的URL。
*/

type Paras struct {
	Id     string
	Name   string
	Phone  []string
	age    int
	salary float64
}

func pack(p interface{}) (url.URL, error) {

	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Struct {
		return url.URL{}, errors.New("p should be a struct")
	}
	vals := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(v.Type().Field(i).Name)
		}
		switch fieldInfo.Type.Kind() {
		case reflect.String:
			fallthrough
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fallthrough
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fallthrough
		case reflect.Float32, reflect.Float64:
			fallthrough
		case reflect.Bool:
			vals.Add(name, fmt.Sprintf("%v", v.Field(i)))
		case reflect.Slice, reflect.Array:
			list := v.Field(i).Interface().([]string)
			for j := 0; j < len(list); j++ {
				vals.Add(name, fmt.Sprintf("%v", list[j]))
			}
		}
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}

func main() {
	var p Paras = Paras{Id: "123", Name: "liusan", Phone: []string{"18100009999", "18611110000"}, age: 45, salary: 8865.5}
	vals, err := pack(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vals.RawQuery)
	//age=45&id=123&name=liusan&phone=18100009999&phone=18611110000&salary=8865.5
}
