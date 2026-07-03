# hard / problem2 — Expression Evaluator (recursive descent)

Theme: **recursion** (recursive-descent parsing). One problem (hard tier). Write
it in `main.go`; `go run .` to eyeball, `go test` to check.

**Allowed packages:** the entire Go standard library.

---

## Evaluate

```go
func Evaluate(expr string) (int, bool)
```

Parse and evaluate an integer arithmetic expression, returning `(value, true)`
on success or `(0, false)` if the expression is malformed.

Grammar / rules:
- **Operands** are non-negative integer literals (one or more digits, multi-digit
  allowed, e.g. `100`). There are **no unary operators** — every `+ - * /` is
  binary, so `"-3"` and `"+3"` are malformed.
- **Operators:** `+`, `-`, `*`, `/`. `*` and `/` bind **tighter** than `+` and
  `-`. Operators of equal precedence are **left-associative**
  (`20 - 5 - 3 == 12`).
- **Parentheses** `(` `)` group sub-expressions and override precedence.
- **Spaces** may appear anywhere and are ignored.
- **Division** is integer division (truncating toward zero, like Go's `/`).
  Division by zero makes the whole expression invalid → `(0, false)`.
- A subtraction may produce a **negative** result (`"3-5" == -2`); that is valid.
- Anything malformed returns `(0, false)`: an empty/blank string, an unknown
  character, a missing operand (`"1+"`, `"*3"`), two operands in a row (`"1 2"`),
  empty parentheses `"()"`, or unbalanced parentheses (`"(1+2"`).

Examples:
```
Evaluate("1+2")         // 3, true
Evaluate("2+3*4")       // 14, true
Evaluate("(2+3)*4")     // 20, true
Evaluate("2*(3+4)-5")   // 9, true
Evaluate("10/3")        // 3, true    (integer division)
Evaluate("20-5-3")      // 12, true   (left-associative)
Evaluate("3-5")         // -2, true
Evaluate(" 3 + 4 ")     // 7, true
Evaluate("1/0")         // 0, false   (division by zero)
Evaluate("1+")          // 0, false
Evaluate("()")          // 0, false
Evaluate("(1+2")        // 0, false
Evaluate("")            // 0, false
```

Hint: classic recursive-descent with one function per precedence level:
- `expr`  = `term` (`(+|-)` `term`)\*    — handles `+` and `-`
- `term`  = `factor` (`(*|/)` `factor`)\* — handles `*` and `/`
- `factor` = number | `(` `expr` `)`     — recurses back to the top for parens

Track a position into the string as you consume tokens; if you reach the end
with everything consumed and no rule violated, it's valid. Strip spaces first to
simplify tokenizing.
