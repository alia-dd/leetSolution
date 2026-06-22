package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("1  *******************************")
// 	fmt.Println(pancakeSort([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// 	fmt.Println("2  *******************************")
// 	fmt.Println(pancakeSort([]int{1, 2, 3}))
// 	fmt.Println("3  *******************************")
// 	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))

// }

// func pancakeSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return []int{}
// 	}
// 	res := []int{}
// 	lastIndex := len(arr)
// 	for range arr {

// 		maxIndex := findMaxInt(arr[:lastIndex])
// 		// fmt.Println(arr[maxIndex])
// 		fmt.Println(maxIndex, lastIndex)
// 		for maxIndex != lastIndex-1 {

// 			if maxIndex == 0 {
// 				arr = flip(arr[:lastIndex])
// 				res = append(res, len(arr[:lastIndex]))
// 				fmt.Println("coorect", arr)
// 				break
// 			}

// 			arr = append(flip(arr[:maxIndex+1]), arr[maxIndex+1:]...)
// 			res = append(res, len(arr[:maxIndex+1]))
// 			maxIndex = findMaxInt(arr[:lastIndex])
// 			fmt.Println("maxIndex after findMax:", maxIndex)

// 			fmt.Println("/////arr;", arr)
// 		}
// 		lastIndex -= 1
// 		fmt.Println(">>>>>arr;", arr, arr[lastIndex])
// 	}
// 	return res
// }
// func findMaxInt(arr []int) int {
// 	maxind := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > arr[maxind] {
// 			// fmt.Println(" arr[i] > arr[maxind]", arr[i], arr[maxind])
// 			maxind = i
// 		}
// 	}
// 	// fmt.Println("max", arr, maxind)
// 	return maxind
// }
// func flip(arr []int) []int {
// 	fmt.Println("flip", arr)
// 	temp := []int{}
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		temp = append(temp, arr[i])
// 	}
// 	return temp
// }
