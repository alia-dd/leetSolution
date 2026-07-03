# easy / battery4 — Contains, MaxSlice, CountOccurrences

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
This battery is about **scanning a slice** recursively: look at the first
element, then recurse on the rest (`nums[1:]`) and combine.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Task 1 — Contains

```go
func Contains(nums []int, target int) bool
```

Report whether `target` appears anywhere in `nums`.

Examples:
```
Contains([]int{}, 5)         // false
Contains([]int{1, 2, 3}, 2)  // true
Contains([]int{1, 2, 3}, 5)  // false
Contains([]int{5}, 5)        // true
```

Hint: base case — the empty slice contains nothing (`false`). Otherwise the
first element matches, or `Contains(nums[1:], target)`.

---

## Task 2 — MaxSlice

```go
func MaxSlice(nums []int) int
```

Return the largest value in `nums`. Assume the slice is **non-empty**.

Examples:
```
MaxSlice([]int{5})           // 5
MaxSlice([]int{1, 2, 3})     // 3
MaxSlice([]int{3, 1, 2})     // 3
MaxSlice([]int{-5, -2, -9})  // -2
```

Hint: base case — a one-element slice's max is that element. Otherwise it's the
larger of `nums[0]` and `MaxSlice(nums[1:])` (the builtin `max` helps).

---

## Task 3 — CountOccurrences

```go
func CountOccurrences(nums []int, target int) int
```

Return how many times `target` appears in `nums`.

Examples:
```
CountOccurrences([]int{}, 2)              // 0
CountOccurrences([]int{1, 2, 2, 3, 2}, 2) // 3
CountOccurrences([]int{1, 2, 3}, 5)       // 0
CountOccurrences([]int{7, 7, 7}, 7)       // 3
```

Hint: add `1` when `nums[0] == target`, then add `CountOccurrences(nums[1:], target)`.
Base case: the empty slice is `0`.
