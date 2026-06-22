package main

// import (
// 	"fmt"
// )

// // func main() {
// // 	// fmt.Println("1  *******************************")
// // 	// fmt.Println(pancakeSort([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// // 	// // fmt.Println("2  *******************************")
// // 	// // fmt.Println(pancakeSort([]int{1, 2, 3}))
// // 	// // fmt.Println("3  *******************************")
// // 	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))

// // }

// func pancakeSort(arr []int) []int {
// 	if len(arr) == 0 {
// 		return []int{}
// 	}
// 	res := []int{}
// 	length := len(arr) - 1
// 	for i := 0; i < length; i++ {
// 		maxint := 0
// 		for k := 0; k < length-i+1; k++ {
// 			if arr[k] > arr[maxint] {
// 				maxint = k
// 			}
// 			fmt.Println("max", arr[maxint])
// 		}

// 		fmt.Println("last max", arr[maxint], "last index", length-i)
// 		temp := []int{arr[maxint]}
// 		arr = append(temp, append(arr[:maxint], arr[maxint+1:]...)...)
// 		if maxint != 0 {
// 			res = append(res, maxint+1)
// 		}

// 		fmt.Println(">", arr)
// 		arr = append(flip(arr[:length-i+1]), arr[length-i+1:]...)
// 		res = append(res, len(arr[:length-i+1]))
// 		fmt.Println(arr)
// 	}
// 	return res
// }

// func flip(arr []int) []int {
// 	temp := []int{}
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		temp = append(temp, arr[i])
// 	}
// 	return temp
// }
func pancakeSort(arr []int) []int {
	if len(arr) < 2 {
		return []int{}
	}
	res := []int{}
	lastIndex := len(arr)
	for range arr {

		maxIndex := findMaxInt(arr[:lastIndex])
		for maxIndex != lastIndex-1 {

			if maxIndex == 0 {
				arr = flip(arr[:lastIndex])
				res = append(res, len(arr[:lastIndex]))
				break
			}
			if maxIndex >= len(arr)/2 {
				arr = flip(arr[:lastIndex])
				maxIndex = findMaxInt(arr[:lastIndex])
				res = append(res, len(arr[:lastIndex]))
			} else {
				arr = append(flip(arr[:len(arr)/2]), arr[(len(arr)/2):]...)
				maxIndex = findMaxInt(arr[:lastIndex])
				res = append(res, len(arr)/2)
			}
		}
		lastIndex -= 1
	}
	return res
}
func findMaxInt(arr []int) int {
	maxind := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxind] {
			maxind = i
		}
	}
	return maxind
}
func flip(arr []int) []int {
	temp := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		temp = append(temp, arr[i])
	}
	return temp
}
