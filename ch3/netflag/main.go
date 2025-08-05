package main

import (
	"fmt"
	"net"
)

func IsUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}

func TurnDown(v *net.Flags) {
	*v &^= net.FlagUp
}

func SetBroadcast(v *net.Flags) {
	*v |= net.FlagBroadcast
}

func IsCast(v net.Flags) bool {
	return v&(net.FlagBroadcast|net.FlagMulticast) != 0
}

func main() {
	fmt.Println(uint(net.FlagUp))
	fmt.Println(uint(net.FlagBroadcast))
	fmt.Println(uint(net.FlagLoopback))
	fmt.Println(uint(net.FlagPointToPoint))
	fmt.Println(uint(net.FlagMulticast))

	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	fmt.Printf("%b %t\n")

}
