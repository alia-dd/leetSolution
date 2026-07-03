# hard / problem1 — N-Queens (count solutions)

Theme: **recursion** (backtracking). One problem (hard tier). Write it in
`main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## CountNQueens

```go
func CountNQueens(n int) int
```

Return the number of distinct ways to place `n` queens on an `n x n` chessboard
so that **no two queens attack each other** — i.e. no two share a row, a column,
or a diagonal. Assume `n >= 1`.

Reference counts:
```
CountNQueens(1) // 1
CountNQueens(2) // 0
CountNQueens(3) // 0
CountNQueens(4) // 2
CountNQueens(5) // 10
CountNQueens(6) // 4
CountNQueens(7) // 40
CountNQueens(8) // 92
```

Hint: place **one queen per row**, choosing its column, and recurse to the next
row (this guarantees no two queens share a row automatically). Keep track of the
columns and the two diagonal directions already used so you can reject an
attacking placement. When you have placed a queen in every row, you've found one
valid arrangement — count it and backtrack to try other columns.

- A queen at `(row, col)` shares a "↘" diagonal with any queen where
  `row - col` is equal, and a "↙" diagonal where `row + col` is equal.
- Base case: once `row == n`, all rows are filled → that's `1` complete solution.
