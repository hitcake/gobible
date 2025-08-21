package main

import "fmt"

func main() {
	/*
	 4.1
	*/
	{
		var a [3]int
		fmt.Println(a[0])        // the first element
		fmt.Println(a[len(a)-1]) // the last element
		//Print the indices and elements
		for i, v := range a {
			fmt.Printf("%d %d\n", i, v)
		}
		// ignore the index, print the elements only
		for _, v := range a {
			fmt.Printf("%d\n", v)
		}
		/*
			默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0。我们也可以使用数组字面值语法用一组值来初始化数组：
		*/
		var q [3]int = [3]int{1, 2, 3}
		var r [3]int = [3]int{1, 2}
		fmt.Println(q)
		fmt.Println(r[2]) // "0"
		/*
			在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。因此，上面q数组的定义可以简化为
		*/
		{
			q := [...]int{1, 2, 3}
			fmt.Printf("%T\n", q) // "[3]int"
		}
		/*
			数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
		*/
		{
			q := [3]int{1, 2, 3}
			//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
			fmt.Printf("%T\n", q)
		}

		/*
			也可以指定一个索引和对应值列表的方式初始化
		*/
		{
			type Currency int

			const (
				USD Currency = iota // 美元
				EUR                 // 欧元
				GBP                 // 英镑
				RMB                 // 人民币
			)

			symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

			fmt.Println(RMB, symbol[RMB]) // "3 ￥"
			/*
				在这种形式的数组字面值形式中，初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初始值的元素将用零值初始化。例如，
			*/
			r := [...]int{99: -1}
			fmt.Println(r)
		}
		/*
			如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我们可以直接通过==比较运算符来比较两个数组，
			只有当两个数组的所有元素都是相等的时候数组才是相等的。不相等比较运算符!=遵循同样的规则。
		*/
		{
			a := [2]int{1, 2}
			b := [...]int{1, 2}
			c := [2]int{1, 3}
			fmt.Println(a == b, a == c, b == c) // "true false false"
			d := [3]int{1, 2}
			fmt.Println(d)
			//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
		}
		/*
			因为函数参数传递的机制导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，并不能直接修改调用时原始的数组变量。
			在这个方面，Go语言对待数组的方式和其它很多编程语言不同，其它编程语言可能会隐式地将数组作为引用或指针对象传入被调用的函数。
			当然，我们可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。下面的函数用于给[32]byte类型的数组清零：
		*/
		{
			//func zero(ptr *[32]byte) {
			//for i := range ptr {
			//	ptr[i] = 0
			//}
			//}
			//func zero(ptr *[32]byte) {
			//    *ptr = [32]byte{}
			//}
		}
	}
}
