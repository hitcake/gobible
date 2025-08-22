package main

import (
	"fmt"
	"sort"
)

func main() {
	//make 创建map
	ages := make(map[string]int)
	fmt.Println(ages)
	// 直接初始化创建map
	ages = map[string]int{"alice": 31, "charlie": 34}
	fmt.Println(ages)
	// 创建空map
	ages = map[string]int{}
	fmt.Println(ages)
	// nil的map
	var ages2 map[string]int
	fmt.Println(ages2 == nil)
	fmt.Println(len(ages))
	// nil map禁止存入元素
	// ages["carol"]=21

	// 修改
	ages["alice"] = 32
	ages["charlie"]++
	ages["bob"]++
	fmt.Println(ages)

	// 删除
	delete(ages, "bob")
	fmt.Println(ages)

	//判断是否存在
	if age, ok := ages["bob"]; ok {
		fmt.Printf("bob'a age is = %d\n", age)
	} else {
		fmt.Println("bob'a age is not found")
	}

	// 遍历
	for name, age := range ages {
		fmt.Printf("%s's age is = %d\n", name, age)
	}
	// 固定顺序遍历
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s's age is = %d\n", name, ages[name])
	}

	fmt.Printf("map equal = %t\n", equal(ages, map[string]int{}))

	// slice作为key的方法
	Add([]string{"123", "456", "789"})
	fmt.Printf("count=%d", Count([]string{"123", "456", "789"}))

	seen := make(map[string]struct{}) // set of strings
	s := "zhangsan"
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

var m = map[string]int{}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}
