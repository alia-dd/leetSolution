# medium / problem3 — Flatten Nested Slices

Theme: **recursion**. One problem (medium tier). The data is arbitrarily deep,
so you must **recurse into the nesting**. Write it in `main.go`; `go run .` to
eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Flatten

```go
func Flatten(nested []any) []int
```

Return all the integers contained in `nested`, at any depth, collected into a
single flat `[]int` in **left-to-right, depth-first** order.

The input is a slice where **each element is either**:
- an `int`, or
- a nested `[]any` (which itself follows the same rule, to any depth).

Rules:
- Preserve order: an integer that appears earlier (reading left to right, diving
  into nested slices as you reach them) comes earlier in the output.
- Empty input, or input that contains only empty nested slices, returns an empty
  `[]int`.

Examples:
```
Flatten([]any{1, 2, 3})                       // [1 2 3]
Flatten([]any{1, []any{2, 3}, 4})             // [1 2 3 4]
Flatten([]any{[]any{[]any{1}}, 2})            // [1 2]
Flatten([]any{[]any{1, 2}, []any{3, []any{4, 5}}}) // [1 2 3 4 5]
Flatten([]any{})                              // []
```

Hint: walk the elements of `nested`; use a **type switch** on each element:
```go
switch v := elem.(type) {
case int:
    // append v to the result
case []any:
    // recurse: append Flatten(v)... to the result
}
```
The recursion is the `[]any` case — it dives one level deeper and splices the
result back in. A single element that is `[]any` may itself contain more
`[]any`, which is why this must recurse rather than just handle one level.
