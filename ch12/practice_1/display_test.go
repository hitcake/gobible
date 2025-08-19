package practice_1

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func TestDisplay(t *testing.T) {
	salary := make(map[Person]int)
	zhangsan := Person{Name: "zhangsan", Age: 18}
	lisi := Person{Name: "lisi", Age: 18}
	salary[zhangsan] = 4300
	salary[lisi] = 4500
	for key, value := range salary {
		fmt.Printf("%s->%d\n", key, value)
	}
	//Display("salary", salary)

	foo := make(map[[2]string]int)
	foo[[2]string{"a", "b"}] = 1
	foo[[2]string{"c", "d"}] = 2
	Display("foo", foo)
}
