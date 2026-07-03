# medium / problem2 — String Permutations

Theme: **recursion**. One problem (medium tier). Solve it **recursively**.
Write it in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library. Assume ASCII strings.

---

## Permutations

```go
func Permutations(s string) []string
```

Return **every ordering** of the characters of `s`.

Rules:
- A string of length `n` has exactly `n!` permutations — return all of them.
- **Do not deduplicate.** If `s` has repeated characters, permutations that
  look identical are still returned separately (so the length is always `n!`).
  For example `Permutations("aab")` returns 6 strings, not 3.
- The empty string has exactly one permutation (itself): `Permutations("")`
  returns `[""]`.
- **Order does not matter** — the tests sort the result before comparing, so you
  may produce the permutations in whatever order your recursion naturally does.

Examples (shown in one possible order):
```
Permutations("")    // [""]
Permutations("a")   // ["a"]
Permutations("ab")  // ["ab", "ba"]
Permutations("abc") // ["abc", "acb", "bac", "bca", "cab", "cba"]
```

Hint: the standard shape — for each position `i` in `s`, take `s[i]` as the
first character, then recurse on the string with that character removed
(`s[:i] + s[i+1:]`), and prepend `s[i]` to each permutation that comes back.
Base case: a string of length `0` or `1` has just one permutation — itself
(`[]string{s}`).
