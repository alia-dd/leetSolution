package main

import "fmt"

type Student struct {
	Name string
	GPA  float64
}

func mergeSort(students []Student) []Student {
	if len(students) < 2 {
		return students
	}
	fmt.Println(students)
	firstHalve := mergeSort(students[:len(students)/2])
	secontHalve := mergeSort(students[(len(students) / 2):])
	fmt.Println("firstHalve", firstHalve)
	fmt.Println("secontHalve", secontHalve)
	return merge(firstHalve, secontHalve)
}

func merge(first, second []Student) []Student {
	findex := 0
	sindex := 0
	merged := []Student{}

	for findex < len(first) && sindex < len(second) {

		if first[findex].GPA > second[sindex].GPA {
			merged = append(merged, first[findex])
			findex++
		} else if first[findex].GPA == second[sindex].GPA {
			if first[findex].Name < second[sindex].Name {

				merged = append(merged, first[findex])
				findex++
			} else {
				merged = append(merged, second[sindex])
				sindex++
			}
		} else {
			merged = append(merged, second[sindex])
			sindex++
		}
	}
	merged = append(merged, first[findex:]...)
	merged = append(merged, second[sindex:]...)

	return merged
}

// func main() {
// 	students := []Student{
// 		{"Alice", 3.8},
// 		{"Bob", 3.5},
// 		{"Carol", 3.8},
// 		{"Dave", 3.2},
// 		{"Eve", 3.5},
// 	}

// 	sorted := mergeSort(students)
// 	for _, s := range sorted {
// 		fmt.Printf("%s: %.1f\n", s.Name, s.GPA)
// 	}
// }
