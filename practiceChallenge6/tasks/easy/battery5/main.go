package main

import "fmt"

// Repeat — see README, Task 1. Must be recursive (no loops).
func Repeat(s string, n int) string {
	// TODO: implement
	if n < 1 {
		return ""
	}
	if n == 1 {
		return s
	}

	return s + Repeat(s, n-1)
}

// CountChar — see README, Task 2. Must be recursive (no loops).
func CountChar(s string, target byte) int {
	// TODO: implement
	if len(s) < 1 {
		return 0
	}
	res := CountChar(s[1:], target)
	if s[0] == target {
		res += 1
	}
	return res
}

// IsSorted — see README, Task 3. Must be recursive (no loops).
func IsSorted(nums []int) bool {
	// TODO: implement
	if len(nums) < 2 {
		return true
	}
	if len(nums) == 2 {
		return nums[0] <= nums[1]
	}
	if IsSorted(nums[1:]) && nums[0] <= nums[1] {
		return true
	}
	return false
}

func main() {
	// Quick eyeball runs — adjust freely while solving.
	fmt.Println(Repeat("ab", 3))
	fmt.Println(CountChar("hello", 'l'))
	fmt.Println(IsSorted([]int{1, 2, 2, 3}))
}
