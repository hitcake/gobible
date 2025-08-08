package main

import "fmt"

func reverse(prt *[]int) {
	l := len(*prt)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		(*prt)[i], (*prt)[j] = (*prt)[j], (*prt)[i]
	}

}

func main() {
	prt := &[]int{1, 2, 3, 4, 5}
	reverse(prt)
	fmt.Println(*prt)
}
