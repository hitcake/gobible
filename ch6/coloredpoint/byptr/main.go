package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}
