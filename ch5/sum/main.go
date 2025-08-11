package main

import "fmt"

func sum(vals ...int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum
}
func main() {
	fmt.Println("Sum:", sum())
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}
