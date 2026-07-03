# medium / problem5 — Recursive Binary Search

Theme: **recursion** (divide & conquer). One problem (medium tier). Write it in
`main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## BinarySearch

```go
func BinarySearch(sorted []int, target int) int
```

`sorted` is a slice of integers in **ascending order with no duplicates**. Return
the **index** of `target` in `sorted`, or `-1` if it is not present. Implement
binary search **recursively**.

Examples:
```
BinarySearch([]int{1, 3, 5, 7, 9}, 5)  // 2
BinarySearch([]int{1, 3, 5, 7, 9}, 1)  // 0
BinarySearch([]int{1, 3, 5, 7, 9}, 9)  // 4
BinarySearch([]int{1, 3, 5, 7, 9}, 4)  // -1
BinarySearch([]int{}, 5)               // -1
BinarySearch([]int{42}, 42)            // 0
```

Hint: look at the middle element. If it equals `target`, you're done; if `target`
is smaller, recurse on the left half; if larger, recurse on the right half.

The tidiest way to return an **absolute** index is to carry `lo` and `hi` bounds
in a helper (`search(sorted, target, lo, hi)`) and compute
`mid = (lo + hi) / 2`, rather than reslicing. If you do reslice with
`sorted[mid+1:]`, remember to add the offset back to whatever index comes up.
Base case: an empty range (`lo > hi`) means "not found" → `-1`.
