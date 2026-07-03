# easy / battery1 — Factorial, SumDigits, Power

Theme: **recursion**. Each of these must be solved **recursively** — no `for`
or `while` loops. The point is to practise expressing a problem as a smaller
version of itself plus a base case.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library (you won't need much).
Assume every input is within `int` range.

---

## Task 1 — Factorial

```go
func Factorial(n int) int
```

Return `n!` = `1 * 2 * ... * n`. Assume `0 <= n <= 20`.

Rules:
- `Factorial(0)` is `1` (the base case).

Examples:
```
Factorial(0)  // 1
Factorial(1)  // 1
Factorial(5)  // 120
Factorial(10) // 3628800
```

Hint: `n! = n * (n-1)!`, and the recursion bottoms out at `n == 0` (or `n <= 1`).

---

## Task 2 — SumDigits

```go
func SumDigits(n int) int
```

Return the sum of the decimal digits of a non-negative integer `n`. Assume
`n >= 0`.

Examples:
```
SumDigits(0)     // 0
SumDigits(7)     // 7
SumDigits(123)   // 6
SumDigits(99999) // 45
```

Hint: the last digit is `n % 10`, and the rest of the number is `n / 10`.
Base case: a single digit (`n < 10`) is its own sum.

---

## Task 3 — Power

```go
func Power(base, exp int) int
```

Return `base` raised to the power `exp`. Assume `exp >= 0`.

Rules:
- `Power(base, 0)` is `1` for any base (including `Power(0, 0) == 1`).

Examples:
```
Power(2, 0)  // 1
Power(2, 10) // 1024
Power(5, 3)  // 125
Power(-2, 3) // -8
```

Hint: `base^exp = base * base^(exp-1)`, bottoming out at `exp == 0` → `1`.
