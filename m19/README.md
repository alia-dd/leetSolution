# Go Output Formatting Challenges

Three challenges, each in its own package with a `_test.go` file using
table-driven tests (LeetCode-style: a slice of {input, expected output}
cases run through a loop with `t.Run`).

## Structure

```
go-challenges/
├── 01_table/
│   ├── table.go        <- implement FormatTable here
│   └── table_test.go   <- do not edit; run against your implementation
└── 02_minisprintf/
    ├── minisprintf.go      <- implement MySprintf here
    └── minisprintf_test.go <- do not edit; run against your implementation
```

## Challenges

1. **01_table** — Given a slice of `Product` structs, format an aligned
   table where each column's width is derived from its longest entry
   (header or data), not hardcoded. Tests column padding, alignment
   direction (left vs right), and dynamic width calculation.

2. **02_minisprintf** — Implement a simplified `Sprintf` supporting `%d`,
   `%s`, `%f`, `%.Nf`, and `%%`, including Go's real error-string behavior
   for missing args (`%!d(MISSING)`) and type mismatches
   (`%!d(string=oops)`). This is the most involved one: parsing, variadic
   args, and `interface{}` type-switching.

Note: the original plan included a fifth challenge — "right-align a column
of numbers to the widest entry" — but that's effectively a subset of
challenge 1 (which already computes per-column widths dynamically), so it
was dropped to avoid redundancy.

## Running the tests

```bash
cd 01_table && go test -v
cd ../02_minisprintf && go test -v
```

Each will fail until you fill in the `TODO` in the corresponding `.go`
file (not the `_test.go` file).

## A note on verification

I don't have a Go toolchain in my current environment, so I verified the
table-formatting expected outputs and the Sprintf logic by hand-simulating
the padding/parsing logic in Python rather than compiling real Go. The
logic should be sound, but if you hit a mismatch between your correct
implementation and a test's expected string, it's worth checking the test
expectation itself before assuming your code is wrong — particularly
around floating point edge cases like values that don't round cleanly
(e.g. `1.005`), which I deliberately avoided in these tests for that
reason.
