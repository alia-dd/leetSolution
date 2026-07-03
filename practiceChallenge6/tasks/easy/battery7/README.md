# easy / battery7 — ReverseString, FirstIndex, AllEven

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
More string and slice recursion: rebuild a string in reverse, search for a
position, and check a property of every element.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## Task 1 — ReverseString

```go
func ReverseString(s string) string
```

Return `s` with its characters in reverse order.

Examples:
```
ReverseString("")      // ""
ReverseString("a")     // "a"
ReverseString("abc")   // "cba"
ReverseString("hello") // "olleh"
```

Hint: base case — a string of length 0 or 1 is its own reverse. Otherwise
`ReverseString(s[1:]) + s[:1]` (reverse the tail, then stick the first character
on the end).

---

## Task 2 — FirstIndex

```go
func FirstIndex(nums []int, target int) int
```

Return the index of the **first** occurrence of `target` in `nums`, or `-1` if it
is not present.

Examples:
```
FirstIndex([]int{1, 2, 3}, 2)    // 1
FirstIndex([]int{1, 2, 3}, 5)    // -1
FirstIndex([]int{}, 5)           // -1
FirstIndex([]int{5}, 5)          // 0
FirstIndex([]int{1, 2, 2}, 2)    // 1  (first match)
```

Hint: if the slice is empty → `-1`. If `nums[0] == target` → `0`. Otherwise
search the rest: `r := FirstIndex(nums[1:], target)`; if `r == -1` return `-1`,
else return `r + 1` (because you dropped one element off the front).

---

## Task 3 — AllEven

```go
func AllEven(nums []int) bool
```

Report whether **every** element of `nums` is even. An empty slice counts as
`true` (there is no odd element).

Examples:
```
AllEven([]int{})        // true
AllEven([]int{2, 4, 6}) // true
AllEven([]int{2, 3, 4}) // false
AllEven([]int{1})       // false
AllEven([]int{0})       // true  (0 is even)
```

Hint: base case — the empty slice is `true`. Otherwise `nums[0]` must be even
**and** `AllEven(nums[1:])` must hold.
