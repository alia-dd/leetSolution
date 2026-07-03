# easy / battery2 — CountVowels, IsPalindrome, SumSlice

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
These move from numbers (battery1) to **strings and slices**: the recursive
step usually peels one element off the end (or both ends) and recurses on the
rest.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## Task 1 — CountVowels

```go
func CountVowels(s string) int
```

Count the vowels (`a e i o u`, both lower and upper case) in `s`.

Examples:
```
CountVowels("")          // 0
CountVowels("xyz")       // 0
CountVowels("hello")     // 2
CountVowels("AEIOU")     // 5
CountVowels("Recursion") // 4
```

Hint: look at the first byte `s[0]`; add 1 if it's a vowel, then recurse on the
rest `s[1:]`. Base case: the empty string is `0`.

---

## Task 2 — IsPalindrome

```go
func IsPalindrome(s string) bool
```

Report whether `s` reads the same forwards and backwards. Comparison is
**case-sensitive** (`"Anna"` is not a palindrome, `A` != `a`).

Examples:
```
IsPalindrome("")        // true
IsPalindrome("a")       // true
IsPalindrome("abba")    // true
IsPalindrome("racecar") // true
IsPalindrome("hello")   // false
```

Hint: compare the first and last characters; if they match, recurse on the
middle `s[1:len(s)-1]`. Base case: length 0 or 1 is `true`.

---

## Task 3 — SumSlice

```go
func SumSlice(nums []int) int
```

Return the sum of all integers in the slice.

Examples:
```
SumSlice([]int{})           // 0
SumSlice([]int{5})          // 5
SumSlice([]int{1, 2, 3})    // 6
SumSlice([]int{10, 20, 30}) // 60
```

Hint: `nums[0] + SumSlice(nums[1:])`. Base case: the empty slice is `0`.
