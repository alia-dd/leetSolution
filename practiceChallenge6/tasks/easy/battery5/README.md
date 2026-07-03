# easy / battery5 — Repeat, CountChar, IsSorted

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
More string and slice recursion: **build up** a result (Repeat), **tally** while
scanning (CountChar), and **check a property of neighbours** (IsSorted).

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## Task 1 — Repeat

```go
func Repeat(s string, n int) string
```

Return `s` concatenated `n` times. Assume `n >= 0`.

Examples:
```
Repeat("ab", 3) // "ababab"
Repeat("x", 0)  // ""
Repeat("", 5)   // ""
Repeat("go", 1) // "go"
```

Hint: base case — `n == 0` returns `""`. Otherwise `s + Repeat(s, n-1)`.

---

## Task 2 — CountChar

```go
func CountChar(s string, target byte) int
```

Return how many times the byte `target` appears in `s`.

Examples:
```
CountChar("hello", 'l') // 2
CountChar("", 'a')      // 0
CountChar("aaa", 'a')   // 3
CountChar("abc", 'z')   // 0
```

Hint: add `1` when `s[0] == target`, then recurse on `s[1:]`. Base case: the
empty string is `0`.

---

## Task 3 — IsSorted

```go
func IsSorted(nums []int) bool
```

Report whether `nums` is sorted in **non-decreasing** order (equal neighbours
are allowed). An empty slice and a single element count as sorted.

Examples:
```
IsSorted([]int{})           // true
IsSorted([]int{1})          // true
IsSorted([]int{1, 2, 2, 3}) // true
IsSorted([]int{3, 1, 2})    // false
IsSorted([]int{5, 4})       // false
```

Hint: base case — length 0 or 1 is `true`. Otherwise `nums[0] <= nums[1]` **and**
`IsSorted(nums[1:])`.
