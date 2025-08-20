package main

import (
	"flag"
	"fmt"
	"gobible/ch2/tempconf"
)

type celsiusFlag struct {
	tempconf.Celsius
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = tempconf.Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = tempconf.FToC(tempconf.Fahrenheit(value))
		return nil
		// 练习7.6
	case "K", "°K":
		c.Celsius = tempconf.KToC(tempconf.Kelvin(value))
		return nil

	}
	return fmt.Errorf("invalid temperature %s", s)
}

func CelsiusFlag(name string, value tempconf.Celsius, usage string) *tempconf.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
