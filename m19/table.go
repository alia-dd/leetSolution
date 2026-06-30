package m19

import (
	"fmt"
	"strconv"
)

// Challenge 1: Aligned Table Formatter
//
// Given a slice of Product structs, return a formatted table as a string
// where each column is padded to the width of its longest entry (including
// the header). Columns are separated by exactly two spaces.
//
// Column alignment rules:
//   - Name:  left-aligned
//   - Price: right-aligned, always shown with 2 decimal places (no currency
//     symbol, no thousands separators -- that's challenge #2's job)
//   - Qty:   right-aligned
//
// Header row is "Name", "Price", "Qty" (exact casing), followed by the data
// rows in the same order they appear in the input slice. There is no
// separator line between header and rows, and no trailing newline after the
// last row. Each row (including the header) is joined with "\n".
//
// Example:
//
//	Input:
//	  []Product{
//	    {"Apple", 1.5, 10},
//	    {"Wholemeal Bread", 3.2, 3},
//	    {"Milk", 0.99, 100},
//	  }
//
//	Output (note column widths derived from longest entries):
//	  Name             Price  Qty
//	  Apple             1.50   10
//	  Wholemeal Bread   3.20    3
//	  Milk              0.99  100
//
// Edge cases to handle:
//   - Empty input slice -> return just the header row.
//   - A single row.
//   - Negative prices (e.g. a refund/adjustment row) -- still right-aligned,
//     "-1.50" counts toward the column's width like any other string.
//   - Price values that need rounding to 2 decimals (e.g. 1.005 -> "1.01"
//     or "1.00" depending on floating point representation -- use standard
//     Go rounding via fmt's %.2f).
//   - Very large quantities/prices that make the Qty/Price column wider
//     than its header.
//
// Constraints:
//   - 0 <= len(products) <= 1000
//   - len(p.Name) <= 100, ASCII printable characters only
//   - -1,000,000 <= p.Price <= 1,000,000
//   - 0 <= p.Qty <= 1,000,000

type Product struct {
	Name  string
	Price float64
	Qty   int
}

// FormatTable returns the aligned, padded table described above.
func FormatTable(products []Product) string {
	fmt.Println("----------------------------")
	// TODO: implement
	if products == nil {
		return fmt.Sprintf("%s  %s  %s", "Name", "Price", "Qty")
	}
	line := ""
	lenght, Qlen, Plen := longS(products)

	// fmt.Println("final", lenght, Qlen, Plen)
	line += fmt.Sprintf("%-*s  %*s  %*s", lenght, "Name", Plen, "Price", Qlen, "Qty")
	line += "\n"
	for i, p := range products {
		if i > 0 {
			line += "\n"
		}

		line += fmt.Sprintf("%-*s  %*.2f  %*d", lenght, p.Name, Plen, p.Price, Qlen, p.Qty)

	}
	// fmt.Println(line)
	return line
}
func longS(products []Product) (int, int, int) {
	lenght := 4
	qlen := 3
	plen := 5
	for _, p := range products {

		if len(p.Name) > lenght {
			lenght = len(p.Name)
		}
		if len(strconv.Itoa(p.Qty)) > qlen {
			qlen = len(strconv.Itoa(p.Qty))
		}
		price := len(strconv.FormatFloat(p.Price, 'f', -1, 64))
		fmt.Println(p.Price, "len", price)
		if price > plen {
			fmt.Println(p.Price, "len", price)
			plen = price
		}
	}

	return lenght, qlen, plen
}
