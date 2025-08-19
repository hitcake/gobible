package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Email    string `check:"email"`
	CardNo   string `check:"card"`
	Postcode int    `check:"postcode"`
}

// 邮箱验证的正则表达式
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var cardRegex = regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`)
var postcodeRegex = regexp.MustCompile(`^[0-9]{5}$`)

func emailChecker(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.String {
		return fmt.Errorf("expected a string")
	}
	value := i.(string)
	if !emailRegex.MatchString(value) {
		return fmt.Errorf("%s not a valid email", value)
	}
	return nil
}

func cardChecker(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.String {
		return fmt.Errorf("expected a string")
	}
	value := i.(string)
	if !cardRegex.MatchString(value) {
		return fmt.Errorf("%s not a valid card", value)
	}
	return nil
}

func postcodeChecker(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.String {
		return fmt.Errorf("expected a int")
	}
	value := i.(string)
	if !postcodeRegex.MatchString(value) {
		return fmt.Errorf("%s not a valid postcode", value)
	}
	return nil
}

type Checker func(v interface{}) error

var checkersMap = map[string]Checker{
	"email":    emailChecker,
	"card":     cardChecker,
	"postcode": postcodeChecker,
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	checkers := make(map[string]Checker)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
		ch := tag.Get("check")
		if ch != "" {
			if checker, ok := checkersMap[ch]; ok {
				checkers[name] = checker
			}
		}
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if checker, ok := checkers[name]; ok {
				err := checker(value)
				if err != nil {
					return err
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("Unsupported type: %s", v.Kind())
	}
	return nil
}

func main() {
	req := &http.Request{Form: url.Values{"email": []string{"123.com"}}}
	if err := Unpack(req, &Data{}); err != nil {
		log.Println(err)
	}
	req = &http.Request{Form: url.Values{"email": []string{"123@126.com"}, "cardno": []string{"123456"}}}
	if err := Unpack(req, &Data{}); err != nil {
		log.Println(err)
	}
	req = &http.Request{Form: url.Values{"email": []string{"123@126.com"}, "cardno": []string{"4999999999999123"}, "postcode": []string{"123456"}}}
	if err := Unpack(req, &Data{}); err != nil {
		log.Println(err)
	}
}
