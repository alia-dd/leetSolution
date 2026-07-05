package main

import "fmt"

// ReverseString — see README, Task 1. Must be recursive (no loops).
func ReverseString(s string) string {
	// TODO: implement
	if len(s) <= 1 {
		return s
	}
	return ReverseString(s[1:]) + s[:1]
}

// FirstIndex — see README, Task 2. Must be recursive (no loops).
func FirstIndex(nums []int, target int) int {
	// TODO: implement
	if len(nums) == 0 {
		return -1
	}
	if nums[len(nums)-1] == target {
		res := FirstIndex(nums[:len(nums)-1], target)
		if res >= 0 {
			return res
		}
		return len(nums) - 1
	}
	res := FirstIndex(nums[:len(nums)-1], target)
	return res
}

// AllEven — see README, Task 3. Must be recursive (no loops).
func AllEven(nums []int) bool {
	// TODO: implement
	if len(nums) == 0 {
		return true
	}
	if AllEven(nums[1:]) && (nums[0]%2 == 0) {
		return true
	}
	return false
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(ReverseString("hello"))
	// fmt.Println(FirstIndex([]int{1, 2, 3, 6, 7, 4}, 9))
	fmt.Println(FirstIndex([]int{1, 3, 2, 2}, 2))
	fmt.Println(AllEven([]int{2, 4, 6}))
}
