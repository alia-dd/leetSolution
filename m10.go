package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println("*******************************")
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// 	fmt.Println("*******************************")
// 	// fmt.Println(convertDateToBinary("26th May 1960"))
// }
func maxProfit(prices []int) int {
	totalpr := 0
	// jump := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			totalpr += prices[i+1] - prices[i]
		}
	}
	return totalpr
}
