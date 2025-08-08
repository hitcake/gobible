package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months[1:])
	fmt.Println(cap(months))
	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:13]
	fmt.Printf("Q1=%s,cap=%d\n", Q1, cap(Q1))
	fmt.Printf("Q2=%s,cap=%d\n", Q2, cap(Q2))
	fmt.Printf("Q3=%s,cap=%d\n", Q3, cap(Q3))
	fmt.Printf("Q4=%s,cap=%d\n", Q4, cap(Q4))
	/**
	Q1=[January February March],cap=12
	Q2=[April May June],cap=9
	Q3=[July August September],cap=6
	Q4=[October November December],cap=3
	*/
	summer := months[6:9]
	fmt.Printf("summer=%s,cap=%d\n", summer, cap(summer))
	summer = summer[:5]
	fmt.Printf("summer=%s,cap=%d\n", summer, cap(summer))
	summer = append(summer, "123", "456", "789", "1000")
	fmt.Printf("summer=%s,cap=%d\n", summer, cap(summer))
	fmt.Printf("months=%s,cap=%d\n", months, cap(months))

}
