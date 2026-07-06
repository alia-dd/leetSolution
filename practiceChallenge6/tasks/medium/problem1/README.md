# medium / problem1 — Tower of Hanoi

Theme: **recursion**. One problem (medium tier). Solve it **recursively** — the
whole point of Hanoi is that the solution is defined in terms of smaller
versions of itself. Write it in `main.go`; `go run .` to eyeball, `go test` to
check.

**Allowed packages:** the entire Go standard library.

---

## Hanoi

```go
func Hanoi(n int, from, to, via string) []string
```

Return the ordered list of moves that transfers a stack of `n` disks from peg
`from` to peg `to`, using peg `via` as the spare, following the rules of the
Tower of Hanoi:

- move exactly one disk at a time (the top disk of a peg),
- noever place a larger disk on tp of a smaller one.

Each move is a string `"X->Y"` meaning "move the top disk from peg `X` to peg
`Y`". Assume `n >= 0` and that `from`, `to`, `via` are three distinct peg names.

Rules:
- `n == 0` → no moves at all (return an empty slice).
- The number of moves for `n` disks is always `2^n - 1`.

Examples:
```
Hanoi(0, "A", "C", "B") // []
Hanoi(1, "A", "C", "B") // ["A->C"]
Hanoi(2, "A", "C", "B") // ["A->B", "A->C", "B->C"]
Hanoi(3, "A", "C", "B") // ["A->C","A->B","C->B","A->C","B->A","B->C","A->C"]
```

Hint: the classic three-step recursion —
1. move the top `n-1` disks from `from` to `via` (using `to` as spare):
   `Hanoi(n-1, from, via, to)`
2. move the largest disk `from -> to` (one move, `"from->to"`)
3. move those `n-1` disks from `via` to `to` (using `from` as spare):
   `Hanoi(n-1, via, to, from)`

Concatenate the three parts in that order. Base case: `n == 0` returns no moves.
Build move strings with `from + "->" + to` (or `fmt.Sprintf`).
