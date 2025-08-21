package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(x + y)

	s := "你好，世界"
	prefix := "你好"
	fmt.Println(HasPrefix(s, prefix))

	for i, r := range s {
		fmt.Printf("%d: %c: %d\n", i, r, r)
	}
	fmt.Println(string(65))    // "A", not "65"
	fmt.Println(string(30028)) // "界"

	b := []byte(s)
	fmt.Println(string(b))

	{
		x := 123
		y := fmt.Sprintf("%d", x)
		fmt.Println(y, strconv.Itoa(x)) // "123 123"

		fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
		//fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多
		s := fmt.Sprintf("x=%b", x) // "x=1111011"
		fmt.Println(s)

		x, err := strconv.Atoi("123") // x is an int
		if err != nil {
			fmt.Println(err)
		}
		z, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(z)
	}

	{
		// math.Pi无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方：
		var x float32 = math.Pi
		var y float64 = math.Pi
		var z complex128 = math.Pi
		fmt.Printf("x: %f, y: %f, z: %.2f\n", x, y, z)
	}
	{
		/*
			math.Pi被确定为特定类型，比如float64，那么结果精度可能会不一样，同时对于需要float32或complex128类型值的地方则会强制需要一个明确的类型转换：
		*/
		const Pi64 float64 = math.Pi

		var x float32 = float32(Pi64)
		var y float64 = Pi64
		var z complex128 = complex128(Pi64)
		fmt.Printf("x: %f, y: %f, z: %.2f\n", x, y, z)
	}
	{
		/*
			除法运算符/会根据操作数的类型生成对应类型的结果。因此，不同写法的常量除法表达式可能对应不同的结果：
		*/
		var f float64 = 212
		fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
		fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
		fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
	}

	{
		/*
			只有常量可以是无类型的。当一个无类型的常量被赋值给一个变量的时候，就像下面的第一行语句，或者出现在有明确类型的变量声明的右边，如下面的其余三行语句，无类型的常量将会被隐式转换为对应的类型，如果转换合法的话
		*/
		{
			var f float64 = 3 + 0i // untyped complex -> float64
			f = 2                  // untyped integer -> float64
			f = 1e123              // untyped floating-point -> float64
			f = 'a'                // untyped rune -> float64
			fmt.Println(f)
		}
		{
			var f float64 = float64(3 + 0i)
			f = float64(2)
			f = float64(1e123)
			f = float64('a')
			fmt.Println(f)
		}
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
