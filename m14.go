package main

import "fmt"

// func main() {
// 	// fmt.Println("1  *******************************")
// 	sortColors([]int{0, 2, 1})
// 	fmt.Println("2  *******************************")
// 	sortColors([]int{2, 0, 2, 1, 1, 0})
// 	fmt.Println("3  *******************************")
// 	sortColors([]int{2, 0, 2, 1, 1, 0})

// }
func sortColors(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 && nums[i] != 1 && nums[i] != 2 {
			return
		}
		small := i
		for k := i + 1; k < len(nums); k++ {
			if nums[k] < nums[small] {
				small = k
			}
		}

		nums[i], nums[small] = nums[small], nums[i]
		fmt.Println(nums)

	}
}
