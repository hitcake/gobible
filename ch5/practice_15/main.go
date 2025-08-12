package main

import (
	"fmt"
	"os"
)

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no values")
	}
	max := vals[0]
	for i := 1; i < len(vals); i++ {
		if max < vals[i] {
			max = vals[i]
		}
	}
	return max, nil
}
func max0(first int, vals ...int) int {
	max := first
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}
func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no values")
	}
	min := vals[0]
	for i := 1; i < len(vals); i++ {
		if min > vals[i] {
			min = vals[i]
		}
	}
	return min, nil
}

func min0(first int, vals ...int) int {
	min := first
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}

func main() {
	minValue, err := min(1, 2, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(minValue)
	}

	fmt.Println(min0(1, 2, 3))
}
