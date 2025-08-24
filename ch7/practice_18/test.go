package main

import "fmt"

func main() {
	map1 := map[string]string{
		"a": "b",
		"b": "c",
	}
	map2 := map[string]string{
		"a": "1",
	}
	fmt.Println(match(map1, map2))
}

func match(x, y map[string]string) bool {
	for k, v := range y {
		if xv, ok := x[k]; ok && xv == v {
			continue
		}
		return false
	}
	return true
}
