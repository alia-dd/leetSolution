package m19

import "testing"

func TestFormatTable(t *testing.T) {
	tests := []struct {
		name     string
		products []Product
		want     string
	}{
		{
			name:     "empty input returns header only",
			products: []Product{},
			want:     "Name  Price  Qty",
		},
		{
			name: "single row",
			products: []Product{
				{"Apple", 1.5, 10},
			},
			want: "Name   Price  Qty\n" +
				"Apple   1.50   10",
		},
		{
			name: "example from problem statement",
			products: []Product{
				{"Apple", 1.5, 10},
				{"Wholemeal Bread", 3.2, 3},
				{"Milk", 0.99, 100},
			},
			want: "Name             Price  Qty\n" +
				"Apple             1.50   10\n" +
				"Wholemeal Bread   3.20    3\n" +
				"Milk              0.99  100",
		},
		{
			name: "negative price widens price column",
			products: []Product{
				{"Refund", -12.50, 1},
				{"Pen", 0.99, 5},
			},
			want: "Name    Price  Qty\n" +
				"Refund  -12.50    1\n" +
				"Pen       0.99    5",
		},
		{
			name: "large quantity widens qty column beyond header",
			products: []Product{
				{"Rice", 0.5, 1000000},
			},
			want: "Name  Price      Qty\n" +
				"Rice   0.50  1000000",
		},
		{
			name: "name longer than header widens name column",
			products: []Product{
				{"A", 1.0, 1},
				{"VeryLongProductNameHere", 2.0, 2},
			},
			want: "Name                     Price  Qty\n" +
				"A                         1.00    1\n" +
				"VeryLongProductNameHere   2.00    2",
		},
		{
			name: "zero quantity and zero price",
			products: []Product{
				{"Free Sample", 0.0, 0},
			},
			want: "Name         Price  Qty\n" +
				"Free Sample   0.00    0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatTable(tc.products)
			if got != tc.want {
				t.Errorf("FormatTable() mismatch\n--- got ---\n%s\n--- want ---\n%s", got, tc.want)
			}
		})
	}
}
