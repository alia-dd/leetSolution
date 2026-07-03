# easy / battery3 — Fibonacci, GCD, CountDown

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
This battery covers three classic recursive shapes: a **two-branch** recursion
(Fibonacci), a **shrinking-argument** recursion (GCD), and **building a slice**
on the way back up (CountDown).

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Task 1 — Fibonacci

```go
func Fibonacci(n int) int
```

Return the `n`-th Fibonacci number, where `F(0) = 0`, `F(1) = 1`, and
`F(n) = F(n-1) + F(n-2)`. Assume `n >= 0` (kept small in the tests).

Examples:
```
Fibonacci(0)  // 0
Fibonacci(1)  // 1
Fibonacci(2)  // 1
Fibonacci(7)  // 13
Fibonacci(10) // 55
```

Hint: two base cases (`n == 0` and `n == 1`), then sum the two smaller calls.

---

## Task 2 — GCD

```go
func GCD(a, b int) int
```

Return the greatest common divisor of `a` and `b` using Euclid's algorithm.
Assume `a, b >= 0` and not both zero.

Examples:
```
GCD(12, 8)  // 4
GCD(48, 36) // 12
GCD(7, 13)  // 1
GCD(10, 0)  // 10
GCD(0, 5)   // 5
```

Hint: `GCD(a, 0) = a` (base case); otherwise `GCD(a, b) = GCD(b, a % b)`.

---

## Task 3 — CountDown

```go
func CountDown(n int) []int
```

Return the slice `[n, n-1, ..., 1, 0]`. Assume `n >= 0`.

Examples:
```
CountDown(0) // [0]
CountDown(1) // [1 0]
CountDown(3) // [3 2 1 0]
CountDown(5) // [5 4 3 2 1 0]
```

Hint: base case `n == 0` returns `[]int{0}`; otherwise prepend `n` to
`CountDown(n-1)` (e.g. `append([]int{n}, CountDown(n-1)...)`).
