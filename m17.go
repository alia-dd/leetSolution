package main

// import "fmt"

// type Student struct {
// 	Name   string
// 	Scores map[string]int // e.g. {"Math": 90, "English": 85}
// }

// func rankStudent(student []Student, subject []string) []Student {
// 	small := 0
// 	for i := 0; i < len(student)-1; i++ {
// 		for k := i + 1; k < len(student); k++ {
// 			for _, sub := range subject {
// 				if student[k].Scores[sub] > student[small].Scores[sub] {
// 					small = k
// 					break
// 				}
// 			}
// 			fmt.Println(student[k], "--------compare-------", student[small])
// 			// if student[k].Scores["Math"] > student[small].Scores["Math"] {
// 			// 	small = k
// 			// } else if student[k].Scores["Math"] == student[small].Scores["Math"] {
// 			// 	fmt.Println(student[k], "--------compare second-------", student[small])
// 			// 	if student[k].Scores["English"] > student[small].Scores["English"] {
// 			// 		small = k
// 			// 	} else if student[k].Scores["English"] == student[small].Scores["English"] {
// 			// 		fmt.Println(student[k], "--------compare thi-------", student[small])
// 			// 		if student[k].Name < student[small].Name {
// 			// 			small = k
// 			// 		}
// 			// 	}
// 			// }

// 		}
// 		fmt.Println(small, student[small])
// 		student[i], student[small] = student[small], student[i]
// 		fmt.Println(student)
// 		small = i + 1
// 	}
// 	return student
// }

// func main() {
// 	students := []Student{
// 		{"Alice", map[string]int{"Math": 90, "English": 85}},
// 		{"Bob", map[string]int{"Math": 90, "English": 92}},
// 		{"Carol", map[string]int{"Math": 78, "English": 85}},
// 		{"Dave", map[string]int{"Math": 90, "English": 85}},
// 	}

// 	fmt.Println("return ", rankStudent(students, []string{"Math", "English"}))

// }
