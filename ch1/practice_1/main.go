package main

import (
	"fmt"
	"os"
)

/*
练习 1.1 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。
*/
func main() {
	fmt.Println(os.Args[0])
}

/*
output:
 practice_1 % go run main.go
/var/folders/x5/6s67syt16dd7_w0r2m2pzbqc0000gn/T/go-build623891956/b001/exe/main
 practice_1 % go build main.go
 practice_1 % ./main
./main
*/
