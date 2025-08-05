package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("%b", x)
	fmt.Println(s)

	x, err := strconv.Atoi("456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
	z, err := strconv.ParseInt("456", 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z)
	}

	fmt.Println(z)
}
