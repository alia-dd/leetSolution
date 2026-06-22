package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("1  *******************************")
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// 	fmt.Println("2  *******************************")
// 	fmt.Println(maxArea([]int{1, 1}))
// 	// fmt.Println("3  *******************************")
// 	// fmt.Println(maxArea([]int{}))
// }

func maxArea(height []int) int {
	maxright, maxleft, maxWater := height[0], height[len(height)-1], 0
	res := 0
	fmt.Println("Before := maxR", maxright, "maxL", maxleft)
	for i, k := 0, len(height)-1; i < k; {

		fmt.Println("data := maxR", height[i], "maxL", height[k])
		if height[i] < height[k] {
			fmt.Println("small")
			if maxright < height[i] {
				maxright = height[i]
			}
			maxWater = (k - i) * (min(height[k], height[i]))
			if maxWater > res {
				res = maxWater
			}
			i++
		} else {
			if maxleft < height[k] {
				maxleft = height[k]
			}
			maxWater = (k - i) * (min(height[k], height[i]))
			if maxWater > res {
				res = maxWater
			}
			k -= 1
		}
		// if maxWater < ((height[i] - 1) * height[k]) {
		// 	fmt.Println((height[i] - 1), "maybe water?? ", height[i]-1*height[k])
		// 	maxWater = ((height[i] - 1) * height[k])
		// }

		fmt.Println("maxR", maxright, "maxL", maxleft, "maxWAter", maxWater)
	}
	// res = maxleft - 1*maxright
	return res
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
