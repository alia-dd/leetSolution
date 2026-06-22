package main

import (
	"fmt"
	"math"
	"time"
)

// func main() {
// 	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
// 	fmt.Println("dddddddddddd")
// 	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))

// }

func findMinDifference(timePoints []string) int {
	if len(timePoints) < 2 || len(timePoints) >= int(2*math.Pow(10, 4)) {
		return 0
	}

	mindiv := math.MaxInt

	for i, t := range timePoints {

		// fmt.Println("time string:..", t)

		timefirst, err := time.Parse("15:04", t)
		if err != nil {
			fmt.Println("f")
			return 0
		}
		fmt.Println("cuttrnt min:", mindiv)
		for k := i + 1; k < len(timePoints); k++ {
			sectime, errsec := time.Parse("15:04", timePoints[k])
			if errsec != nil {
				fmt.Println("f")
				return 0
			}
			diff := int(math.Abs(float64(timefirst.Sub(sectime).Minutes())))
			fmt.Println("diff >>", diff)
			if diff > 720 {
				diff = 1440 - diff
			}
			if diff < mindiv {
				mindiv = diff
			}
		}
	}
	return mindiv
}
