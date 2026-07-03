# medium / problem4 — Subset Sum

Theme: **recursion**. One problem (medium tier). Classic **include-or-exclude**
recursion. Write it in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## SubsetSum

```go
func SubsetSum(nums []int, target int) bool
```

Report whether **some subset** of `nums` adds up exactly to `target`. The subset
may be empty (which sums to `0`), and elements may not be reused.

Rules:
- The empty subset sums to `0`, so any `target == 0` is reachable (`true`).
- Each element is either **included** in the subset or **excluded** — that binary
  choice is the recursion.

Examples:
```
SubsetSum([]int{3, 34, 4, 12, 5, 2}, 9)  // true  (4 + 5, or 3 + 4 + 2)
SubsetSum([]int{3, 34, 4, 12, 5, 2}, 30) // false
SubsetSum([]int{1, 2, 3}, 0)             // true  (empty subset)
SubsetSum([]int{2, 4}, 3)                // false
SubsetSum([]int{5}, 5)                   // true
SubsetSum([]int{}, 0)                    // true
SubsetSum([]int{}, 5)                    // false
```

Hint: consider the first element `nums[0]`. Either it's part of a valid subset
(recurse on `nums[1:]` with `target - nums[0]`) **or** it isn't (recurse on
`nums[1:]` with the same `target`). Return `true` if **either** branch succeeds.
Base cases: `target == 0` → `true`; empty slice with non-zero target → `false`.
