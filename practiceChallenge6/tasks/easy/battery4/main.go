package main

import (
	"fmt"
)

// Contains — see README, Task 1. Must be recursive (no loops).
func Contains(nums []int, target int) bool {
	// TODO: implement
	if len(nums) < 1 {
		return false
	}
	if nums[0] == target || Contains(nums[1:], target) {
		return true
	}
	return false
}

// MaxSlice — see README, Task 2. Must be recursive (no loops). Assume non-empty.
func MaxSlice(nums []int) int {
	// TODO: implement
	if len(nums) == 1 {
		return nums[0]
	}
	res := MaxSlice(nums[1:])
	if nums[0] > res {
		return nums[0]
	}
	return res
}

// CountOccurrences — see README, Task 3. Must be recursive (no loops).
func CountOccurrences(nums []int, target int) int {
	// TODO: implement
	if len(nums) < 1 {
		return 0
	}
	num := CountOccurrences(nums[1:], target)
	if nums[0] == target {
		num += 1
	}
	return num
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(MaxSlice([]int{3, 1, 2}))
	fmt.Println(CountOccurrences([]int{1, 2, 2, 3, 2}, 2))
}
