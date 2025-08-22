package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Addresss  string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Name = "Dilbert"
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(employeeOfTheMonth.ManagerID).Position)
	EmployeeByID(employeeOfTheMonth.ID).Position += " (proactive team player)"
}

/*
EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量
（译注：调用函数返回的是值，并不是一个可取地址的变量）
*/
func EmployeeByID(id int) *Employee {
	return &Employee{ID: id, Name: "Dilbert", Position: "Senior"}
}

/*
结构体成员的输入顺序也有重要的意义。我们也可以将Position成员合并（因为也是字符串类型），或者是交换Name和Address出现的先后顺序，
那样的话就是定义了不同的结构体类型。通常，我们只是将相关的成员写到一起。
type Employee struct {
    ID            int
    Name, Address string
    DoB           time.Time
    Position      string
    Salary        int
    ManagerID     int
}
*/
