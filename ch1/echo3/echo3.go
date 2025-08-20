package main

import (
	"fmt"
	"os"
	"strings"
)

/*
每次循环迭代字符串 s 的内容都会更新。+= 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 s。
s 原来的内容已经不再使用，将在适当时机对它进行垃圾回收。
如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用 strings 包的 Join 函数：
*/
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
