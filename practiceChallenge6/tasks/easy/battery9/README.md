# easy / battery9 — RemoveChar, ConcatSlice, Range

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
The last easy battery: filter characters out of a string, fold a slice of
strings into one, and build a slice of numbers.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## Task 1 — RemoveChar

```go
func RemoveChar(s string, c byte) string
```

Return `s` with **every** occurrence of the byte `c` removed.

Examples:
```
RemoveChar("hello", 'l') // "heo"
RemoveChar("aaa", 'a')   // ""
RemoveChar("abc", 'z')   // "abc"
RemoveChar("", 'a')      // ""
```

Hint: base case — the empty string returns `""`. Otherwise look at `s[0]`: if it
equals `c`, skip it and return `RemoveChar(s[1:], c)`; if not, keep it:
`s[:1] + RemoveChar(s[1:], c)`.

---

## Task 2 — ConcatSlice

```go
func ConcatSlice(strs []string) string
```

Return all the strings joined together into one (no separator).

Examples:
```
ConcatSlice([]string{})                    // ""
ConcatSlice([]string{"a", "b", "c"})       // "abc"
ConcatSlice([]string{"hello", " ", "world"}) // "hello world"
```

Hint: base case — the empty slice returns `""`. Otherwise `strs[0] +
ConcatSlice(strs[1:])`.

---

## Task 3 — Range

```go
func Range(start, end int) []int
```

Return the slice of integers from `start` up to **and including** `end`. If
`start > end`, return an empty slice.

Examples:
```
Range(1, 3) // [1 2 3]
Range(5, 5) // [5]
Range(0, 4) // [0 1 2 3 4]
Range(3, 2) // []
```

Hint: base case — if `start > end`, there are no numbers → return an empty slice.
Otherwise prepend `start` to the rest: `append([]int{start}, Range(start+1, end)...)`.
