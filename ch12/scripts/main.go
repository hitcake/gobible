package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	/*
		函数 reflect.TypeOf 接受任意的 interface{} 类型，并以 reflect.Type 形式返回其动态类型
	*/
	t := reflect.TypeOf(3)
	fmt.Println(t.String()) //int
	fmt.Println(t)          //int
	/*
		reflect.TypeOf 返回的是一个动态类型的接口值，它总是返回具体的类型。
		因此，下面的代码将打印 "*os.File" 而不是 "io.Writer"
	*/
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) //*of.File
	/*
		reflect.Type 接口是满足 fmt.Stringer 接口的。
		因为打印一个接口的动态类型对于调试和日志是有帮助的，
		fmt.Printf 提供了一个缩写 %T 参数，内部使用 reflect.TypeOf 来输出
	*/
	fmt.Printf("%T\n", w)
	/*
		一个 reflect.Value 可以装载任意类型的值。函数 reflect.ValueOf 接受任意的 interface{} 类型，
		并返回一个装载着其动态值的 reflect.Value
	*/
	v := reflect.ValueOf(3)
	fmt.Println(v)          //3
	fmt.Printf("%v\n", v)   //3
	fmt.Println(v.String()) //<int Value>

	/*
		对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type：
	*/
	tv := v.Type()
	fmt.Println(tv.String())
	/*
		reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法。
		它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值
		所不同的是，一个空的接口隐藏了值内部的表示方式和所有方法，因此只有我们
		知道具体的动态类型才能使用类型断言来访问内部的值
	*/

	x := v.Interface()    // an interface{}
	i := x.(int)          // an int
	fmt.Printf("%d\n", i) //3
}
