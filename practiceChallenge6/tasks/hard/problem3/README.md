# hard / problem3 — Word Break

Theme: **recursion**. One problem (hard tier). Write it in `main.go`;
`go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## WordBreak

```go
func WordBreak(s string, dict []string) bool
```

Report whether `s` can be split into a sequence of **one or more** words that
all appear in `dict`. Words from the dictionary may be reused any number of
times, and the whole of `s` must be consumed (no leftover characters).

Rules:
- The empty string `""` is breakable by convention → `true`.
- A word may be used multiple times.
- The split must cover `s` exactly — every character belongs to exactly one word.

Examples:
```
WordBreak("leetcode", []string{"leet", "code"})                       // true
WordBreak("applepenapple", []string{"apple", "pen"})                  // true
WordBreak("catsandog", []string{"cats","dog","sand","and","cat"})     // false
WordBreak("aaaa", []string{"a", "aa"})                                // true
WordBreak("cars", []string{"car", "ca", "rs"})                        // true
WordBreak("a", []string{"b"})                                         // false
```

Hint: try every dictionary word that is a **prefix** of `s`; for each such word,
strip it off the front and recurse on the remaining suffix. Return `true` if any
choice leads to fully consuming the string. Base case: the empty string returns
`true` (you've consumed everything). Beware: `"catsandog"` shows why you must
try *all* matching prefixes and backtrack — an early greedy match can dead-end.
