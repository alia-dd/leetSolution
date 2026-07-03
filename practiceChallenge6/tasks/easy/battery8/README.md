# easy / battery8 — IsEven, IsPowerOfTwo, CollatzSteps

Theme: **recursion**. Solve each one **recursively** — no `for`/`while` loops.
These shrink the number toward a base case in different ways: subtract 2, halve,
or follow the Collatz rule.

Write them all in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Task 1 — IsEven

```go
func IsEven(n int) bool
```

Report whether `n` is even, **without** using `%` or `/`. Assume `n >= 0`.

Examples:
```
IsEven(0) // true
IsEven(1) // false
IsEven(4) // true
IsEven(7) // false
```

Hint: two base cases — `n == 0` is even (`true`), `n == 1` is odd (`false`).
Otherwise the evenness of `n` is the same as that of `n - 2`: `IsEven(n-2)`.

---

## Task 2 — IsPowerOfTwo

```go
func IsPowerOfTwo(n int) bool
```

Report whether `n` is a power of two (`1, 2, 4, 8, 16, ...`). Assume `n >= 1`.

Examples:
```
IsPowerOfTwo(1)  // true   (2^0)
IsPowerOfTwo(2)  // true
IsPowerOfTwo(8)  // true
IsPowerOfTwo(3)  // false
IsPowerOfTwo(12) // false
```

Hint: base case `n == 1` → `true`. If `n` is odd (and greater than 1) it can't be
a power of two → `false`. Otherwise recurse on `n / 2`.

---

## Task 3 — CollatzSteps

```go
func CollatzSteps(n int) int
```

Return how many steps it takes to reach `1` under the Collatz rule: if `n` is
even, go to `n / 2`; if odd, go to `3*n + 1`. Assume `n >= 1`.

Examples:
```
CollatzSteps(1) // 0
CollatzSteps(2) // 1
CollatzSteps(3) // 7
CollatzSteps(6) // 8
```

Hint: base case `n == 1` returns `0` (already there). Otherwise it's `1 +
CollatzSteps(next)`, where `next` is `n/2` if even or `3*n+1` if odd.
