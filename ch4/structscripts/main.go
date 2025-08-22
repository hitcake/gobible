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

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
func ScaleByPtr(p *Point, factor int) *Point {
	return &Point{p.X * factor, p.Y * factor}
}

func main() {
	p := Point{1, 2}
	q := Point{X: 2, Y: 1}
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(Scale(p, 2))
	fmt.Println(ScaleByPtr(&p, 2))

	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	var w Wheel
	w.X = 8
	w.Y = 9
	w.Radius = 5
	w.Spokes = 20
}
