package main

import "fmt"

func removeDuplicatesByNewAlloc(nums []string) []string {
	if len(nums) <= 1 {
		return nums
	}
	var pre string
	ret := make([]string, 0, len(nums))
	for i, v := range nums {
		if i == 0 {
			pre = v
			ret = append(ret, pre)
		} else {
			if pre != v {
				pre = v
				ret = append(ret, v)
			}
		}
	}
	return ret
}
func removeDuplicatesByLocal(nums []string) []string {
	if len(nums) <= 1 {
		return nums
	}
	var pre string
	for i := 0; i < len(nums); {
		if i == 0 {
			pre = nums[i]
			i++
		} else {
			if pre != nums[i] {
				pre = nums[i]
				i++
			} else {
				copy(nums[i:], nums[i+1:])
				nums = nums[:len(nums)-1]
			}
		}
	}
	return nums
}
func main() {
	fmt.Println(removeDuplicatesByLocal([]string{"1", "2", "3", "3", "5"}))
}
