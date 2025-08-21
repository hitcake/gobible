package main

import "fmt"

/*
*

	K>0 向右旋转
	k<0 向左旋转
*/
func rotate(nums []int, k int) []int {
	if k == 0 {
		return nums
	} else {

		result := make([]int, len(nums))
		if k > 0 {
			k = k % len(nums)
			for i, v := range nums {
				if i+k < len(nums) {
					result[i+k] = v
				} else {
					result[i-(len(nums)-k)] = v
				}
			}
		} else {
			k = -k
			for i, v := range nums {
				if i-k >= 0 {
					result[i-k] = v
				} else {
					result[i-k+len(nums)] = v
				}
			}
		}

		return result
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = rotate(nums, -2)
	fmt.Println(nums)
}
