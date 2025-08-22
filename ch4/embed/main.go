package main

import "fmt"

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

//	type Circle struct {
//		X,Y,Radius int
//	}
type Circle struct {
	Point
	Radius int
}

//	type Wheel struct {
//		X,Y,Radius,Spokes int
//	}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle: Circle{Point{2, 3}, 5}, Spokes: 5}
	fmt.Println("%$v\n", w)
	w.X = 42
	fmt.Println("%$v\n", w)
}
