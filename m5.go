package main

import (
	"fmt"
	// "math"
	// "time"
)

// func main() {
// 	fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}}))
// 	fmt.Println("*****************************")
// 	fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}))

// }

func averageWaitingTime(customers [][]int) float64 {
	sub := customers[0][0]
	customerTime := []int{}
	for _, t := range customers {
		if sub < t[0] {
			sub = t[0] + t[1]
			customerTime = append(customerTime, sub-t[0])
			continue
		}
		sub = sub + t[1]
		fmt.Println("t0", t[0], "t1", t[1])
		fmt.Println("add", sub)
		fmt.Println("sub", sub, "-", t[0], "=", sub-t[0])
		customerTime = append(customerTime, sub-t[0])
	}
	fmt.Println("cuT:", customerTime)
	var avarage float64
	for _, av := range customerTime {
		avarage += float64(av)
	}

	avarage = avarage / float64(len(customerTime))
	return avarage
}
