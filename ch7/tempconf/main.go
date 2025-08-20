package main

import (
	"flag"
	"fmt"
	"gobible/ch2/tempconv"
)

type celsiusFlag struct {
	tempconv.Celsius
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
		// 练习7.6
	case "K", "°K":
		c.Celsius = tempconv.KToC(tempconv.Kelvin(value))
		return nil

	}
	return fmt.Errorf("invalid temperature %s", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
