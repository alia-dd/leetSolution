# medium / problem6 — Generate Balanced Parentheses

Theme: **recursion** (backtracking). One problem (medium tier). Write it in
`main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## GenerateParentheses

```go
func GenerateParentheses(n int) []string
```

Return **all** strings of `n` pairs of parentheses that are **well-formed**
(every `(` is matched by a later `)`, and no prefix has more `)` than `(`).
Assume `n >= 0`.

Rules:
- `n == 0` → exactly one string, the empty string: `[""]`.
- The number of results is the `n`-th Catalan number (`1, 1, 2, 5, 14, ...`).
- **Order does not matter** — the tests sort the result before comparing.

Examples (shown in one possible order):
```
GenerateParentheses(0) // [""]
GenerateParentheses(1) // ["()"]
GenerateParentheses(2) // ["(())", "()()"]
GenerateParentheses(3) // ["((()))", "(()())", "(())()", "()(())", "()()()"]
```

Hint: build the string one character at a time, tracking how many `(` and `)`
you've used so far:
- you may add a `(` while `open < n`,
- you may add a `)` while `close < open` (so you never close more than you've
  opened).

Recurse on each allowed choice; when the string reaches length `2*n` it is a
complete, valid arrangement — collect it. Base case: length `2*n` (or
equivalently `open == n && close == n`).
