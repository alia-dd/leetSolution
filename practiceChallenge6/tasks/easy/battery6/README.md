# easy / battery6 — Multiply, CountDigits, SumTo

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
More number recursion: build multiplication out of addition, and peel digits or
counts off one at a time.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Task 1 — Multiply

```go
func Multiply(a, b int) int
```

Return `a * b`, but computed as **repeated addition** (no `*`). Assume `b >= 0`
(`a` may be negative).

Examples:
```
Multiply(3, 4)  // 12
Multiply(5, 0)  // 0
Multiply(0, 7)  // 0
Multiply(-2, 3) // -6
Multiply(7, 1)  // 7
```

Hint: base case `b == 0` returns `0`; otherwise `a + Multiply(a, b-1)`.

---

## Task 2 — CountDigits

```go
func CountDigits(n int) int
```

Return how many decimal digits `n` has. Assume `n >= 0` (`0` has one digit).

Examples:
```
CountDigits(0)     // 1
CountDigits(9)     // 1
CountDigits(10)    // 2
CountDigits(12345) // 5
```

Hint: base case `n < 10` has `1` digit; otherwise `1 + CountDigits(n / 10)`.

---

## Task 3 — SumTo

```go
func SumTo(n int) int
```

Return the sum `1 + 2 + ... + n`. Assume `n >= 0` (`SumTo(0)` is `0`).

Examples:
```
SumTo(0)   // 0
SumTo(1)   // 1
SumTo(5)   // 15
SumTo(100) // 5050
```

Hint: base case `n == 0` returns `0`; otherwise `n + SumTo(n-1)`.
