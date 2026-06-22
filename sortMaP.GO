package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"banana": 3, "apple": 1, "cherry": 2}

	// 1. Extract keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 2. Sort keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// 3. Iterate in sorted order
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}

}
