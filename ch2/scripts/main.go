package main

import (
	"fmt"
	"image"
)

func main() {
	var s string
	fmt.Println(s)

	var i, j, k int
	var b, f, m = true, 2.3, "four"
	fmt.Printf("%d %d %d %t %.2f %s\n", i, j, k, b, f, m)

	n := 100
	var boiling float64 = 100
	var names []string
	var err error
	var p image.Point
	fmt.Printf("%d %.2f %s %T %T\n", n, boiling, names, err, p)

	x := 1
	ptr := &x
	fmt.Println(*ptr)
	*ptr = 2
	fmt.Printf("x=%d\n", x)

	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil) //true false false

	fmt.Println(newf() == newf()) //false

	incr(ptr)
	fmt.Println(*ptr)

	q := new(int)
	fmt.Println(*q)
	*q = 2
	fmt.Println(*q)

	p1 := new(int)
	p2 := new(int)
	fmt.Println(p1 == p2) //false
}

func newf() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

// new 可以作为变量名字
func delta(old, new int) int {
	return old - new
}

var global *int

// 函数执行后x不会回收
func fg() {
	var x int
	x = 1
	global = &x
}

// 函数执行后y会立即回收
func gg() {
	y := new(int)
	*y = 1
}
