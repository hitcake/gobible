package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	/*
		底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。
		这意味着，Celsius和Fahrenheit类型的算术运算行为和底层的float64类型是一样的，
	*/

	fmt.Printf("%g\n", BoilingC-FreezingC) //100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) //100
	//fmt.Printf("%g\n", boilingF-FreezingC)

	var c Celsius
	var f Fahrenheit
	fmt.Println("==========")
	fmt.Println(c == 0) //true
	fmt.Println(f == 0) //true
	//fmt.Println(c==f)
	fmt.Println(c == Celsius(f)) //true
	/*
		注意最后那个语句。尽管看起来像函数调用，但是Celsius(f)是类型转换操作，它并不会改变值，仅仅是改变值的类型而已。测试为真的原因是因为c和f都是零值。
	*/

	c = FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

}
