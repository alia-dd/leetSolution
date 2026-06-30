package main

import (
	"fmt"
)

type Item struct {
	Name  string
	Price int
}
type Students struct {
	Name  string
	Grade int
}
type Row struct {
	Product string
	Amount  int
	Price   float64
}

func main() {
	items := []Item{
		{"Apple", 2},
		{"VeryLongChocolateBarName", 5},
		{"Milk", 3},
	}
	fmt.Println("Task 1........................")
	fmt.Println(FormatReceipt(items))

	students := []Students{
		{"Bob", 8},
		{"Alexander", 10},
		{"Li", 9},
		{"AReallyLongStudentName", 10},
	}
	fmt.Println("\n\nTask 2........................")
	fmt.Println(FormatStudents(students))

	rows := []Row{
		{"Apple", 2, -1.50},
		{"Banana", 10000000000, 0.04},
		{"Ananas", 15, 10000000.416},
		{"SuperLongProductName", 1, 20.00},
	}
	fmt.Println("\n\nTask 3........................")
	fmt.Println(MakeTable(rows))

}

// Task 1 — Receipt (fixed layout + clamp)
// Difficulty: 10 pts

// Implement:
// Given:

// type Item struct {
//     Name  string
//     Price int
// }

// Example input:

// Expected output:

// Item                 Price
// Apple                2€
// VeryLongChocolate    5€
// Milk                 3€
// Total                10€

// Rules:

// Item column width is exactly 20 characters
// If name is longer → cut it
// If shorter → pad with spaces
// Price is right after the column
// No extra spaces at line ends

func FormatReceipt(items []Item) string {
	res := ""
	res += fmt.Sprintf("%-20.20s %s", "Items", "Price")
	res += "\n"
	for i, p := range items {
		// if len(p.Name) > 20 {
		// 	res += fmt.Sprintf("%-20s%d€", p.Name[:17], p.Price)
		// } else {
		res += fmt.Sprintf("%-20.20s%d€", p.Name, p.Price)
		// }

		if i < len(items)-1 {
			res += "\n"
		}
	}
	return res
}

// ***********************************************

// Task 2 — Student Report (dynamic width)
// Difficulty: 30 pts

// Implement:

// Output:

// Name       Grade
// Bob        8
// Alexander  10
// Li         9

// Rules:

// Name column width is based on the longest name
// Header counts as a possible width
// No unnecessary trailing spaces
// Must work with:
// {"AReallyLongStudentName", 10}
// Mental model

// You cannot format immediately.

// You need:

// first loop:
// find width

// second loop:
// print

// Two-pass problem.

func FormatStudents(students []Students) string {
	res := ""
	lenght := lonN(students)
	// format := fmt.Sprintf("%%-%ds%%d", lenght+2)

	// formathead := fmt.Sprintf("%%-%ds  %%s\n", lenght)
	// res += fmt.Sprintf(formathead, "Name", "Grade")
	// fmt.Println(format)
	res += fmt.Sprintf("%-*s Grade\n", lenght, "Name")
	for i, s := range students {
		if i > 0 {
			res += "\n"
		}
		// res += fmt.Sprintf(format, s.Name, s.Grade)

		res += fmt.Sprintf("%-*s %d", lenght, s.Name, s.Grade)

		// 	if i < len(students)-1 {
		// 		res += "\n"
		// 	}
	}
	return res
}

func lonN(students []Students) int {
	lenght := 4
	for _, s := range students {
		if len(s.Name) > lenght {
			lenght = len(s.Name)
		}
	}

	return lenght
}

// ***********************************************
// Task 3 — ASCII Table (hard)
// Difficulty: 70 pts

// Implement:

// Given:

// Output:

// +----------------------+--------+-------+
// | Product              | Amount | Price |
// +----------------------+--------+-------+
// | Apple                | 2      | 1.50  |
// | Banana               | 10     | 0.40  |
// | SuperLongProductName | 1      | 20.00 |
// +----------------------+--------+-------+

// Rules:

// Every column width is dynamic
// Headers affect width
// Numbers are left aligned
// Price always has 2 decimals
// Must work with any input

func MakeTable(rows []Row) string {
	res := ""
	length, lenA, lenP := lonName(rows)
	res += "+" + space(length+2, "-") + "+" + space(lenA+2, "-") + "+" + space(lenP+2, "-") + "+\n"
	res += "| Product " + space((length-len("Product")), " ") + "| Amount" + space((lenA-len("Amount")), " ") + " |" + " Price " + space((lenP-len("Price")), " ") + "|\n"
	res += "+" + space(length+2, "-") + "+" + space(lenA+2, "-") + "+" + space(lenP+2, "-") + "+\n"
	for i, gr := range rows {
		if i > 0 {
			res += " |\n"
		}
		res += "| " + gr.Product + " " + space((length-len(gr.Product)), " ")
		res += "| " + Itoa(gr.Amount) + space(lenA-len(Itoa(gr.Amount)), " ")
		res += " | " + ForFloat(gr.Price, 2) + space((lenP-len(ForFloat(gr.Price, 2))), " ")

	}
	res += " |\n+" + space(length+2, "-") + "+" + space(lenA+2, "-") + "+" + space(lenP+2, "-") + "+"
	return res
}
func lonName(rows []Row) (int, int, int) {
	length, Alen, Plen := 4, 6, 5
	for _, s := range rows {
		if len(s.Product) > length {
			length = len(s.Product)
		}
		if len(Itoa(s.Amount)) > Alen {
			Alen = len(Itoa(s.Amount))
		}
		if len(ForFloat(s.Price, 2)) > Plen {
			Plen = len(ForFloat(s.Price, 2))
		}
	}

	return length, Alen, Plen
}

func space(length int, c string) string {
	res := ""
	for range length {
		res += c
	}
	return res
}

// func ForFloat(num float64, pad int) string {
// 	floatstr := ""
// 	whole := int(num)
// 	frac := num - float64(whole)

// 	floatstr += Itoa(whole) + "."
// 	for range pad {
// 		digit := int(frac * 10)
// 		frac = frac*10 - float64(digit)
// 		floatstr += Itoa(digit)
// 	}

// 	return floatstr
// }

func ForFloat(num float64, pr int) string {
	res := ""
	isn := false

	if num < 0 {
		isn = true
		num *= -1
	}
	whole := int(num)
	frac := num - float64(whole)
	res += Itoa(whole) + "."

	for range pr {
		frac *= 10
	}
	fracint := int(frac + 0.5)
	res += Itoa(fracint)

	if len(Itoa(fracint)) != 2 {
		res = Itoa(whole) + "." + "0" + Itoa(fracint)

	}

	if isn {
		res = "-" + res
	}

	return res

}
