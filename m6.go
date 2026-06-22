package main

import (
	"fmt"
	"math"
)

// import (
// 	"fmt"
// 	"math"
// 	// "time"
// )

// func main() {
// 	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
// 	fmt.Println("*******************************")
// 	fmt.Println(minimumTime([]int{2}, 1))
// 	fmt.Println("*******************************")
// fmt.Println(minimumTime([]int{39,82,69,37,78,14,93,36,66,61,13,58,57,12,70,14,67,75,91,1,34,68,73,50,13,40,81,21,79,12,35,18,71,43,5,50,37,16,15,6,61,7,87,43,27,62,95,45,82,100,15,74,33,95,38,88,91,47,22,82,51,19,10,24,87,38,5,91,10,36,56,86,48,92,10,26,63,2,50,88,9,83,20,42,59,55,8,15,48,25}, 9))

// }
func minimumTime(time []int, totalTrips int) int64 {
	result := make([]int, len(time))
	copy(result, time)
	check := make([]int, len(time))
	count := 0
	if len(time) < 1 || len(time) > int(math.Pow(10, 4)) || totalTrips > 107 {
		return 0
	}
	if len(time) == 1 {
		return int64(time[0])
	}
	// i := 0; i < totalTrips; i++
	for {
		sum := 0
		for _, r := range check {
			sum += r
		}
		fmt.Println("sum", sum)
		if sum >= totalTrips {
			break
		}
		for k, t := range time {

			if t < 1 {
				return 0
			}
			fmt.Println("loop:", result)
			result[k] -= 1
			fmt.Println("minus loop:", result)
			if result[k] == 0 {
				fmt.Println("reached adding", time[k])
				result[k] = time[k]
				check[k] += 1

			}

		}
		count++
		fmt.Println("chck:.........", check)
		fmt.Println("-------------")
	}

	return int64(count)
}

// func minimumTime(time []int, totalTrips int) int64 {
//     lo := int64(1)
//     hi := int64(slices.Min(time)) * int64(totalTrips)

//     for lo < hi {
//         mid := (lo + hi) / 2
//         trips := int64(0)
//         for _, t := range time {
//             trips += mid / int64(t)
//         }
//         if trips >= int64(totalTrips) {
//             hi = mid
//         } else {
//             lo = mid + 1
//         }
//     }
//     return lo
// }
