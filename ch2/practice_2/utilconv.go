package main

import "fmt"

type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

func (m Meter) String() string {
	return fmt.Sprintf("%g meters", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g feet", f)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}
