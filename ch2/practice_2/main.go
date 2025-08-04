package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, _ := strconv.ParseFloat(arg, 64)
		m := Meter(t)
		f := Feet(t)
		k := Kilogram(f)
		p := Pound(f)
		fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), f, FToM(f))
		fmt.Printf("%s = %s, %s = %s\n", k, KToP(k), p, PToK(p))

	}
}
