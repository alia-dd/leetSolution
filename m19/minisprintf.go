package m19

// Challenge 2: MySprintf
//
// Implement MySprintf(format string, args ...interface{}) string, a
// simplified re-implementation of fmt.Sprintf supporting exactly these
// verbs:
//
//	%d  - integer (accepts int)
//	%s  - string  (accepts string)
//	%f  - float, default 6 decimal places, same as fmt's default
//	      (accepts float64)
//	%.Nf - float with N decimal places, where N is a single digit 0-9
//	      (e.g. %.2f, %.0f)
//	%%  - literal percent sign, consumes no argument
//
// Behavior to match fmt.Sprintf for the supported subset:
//   - Verbs are consumed left to right and matched against args in order.
//   - If there are more verbs than args, append "%!d(MISSING)" (or %!s /
//     %!f respectively) in place of the missing value.
//   - If an arg's type doesn't match the verb (e.g. %d given a string),
//     append "%!d(string=<value>)" using the verb that was expected and
//     the Go type name of what was actually passed, followed by =<value>.
//     (Only int/string/float64 inputs will be tested, so you only need to
//     handle those three type names: "int", "string", "float64".)
//   - Extra args beyond the number of verbs are ignored (matches fmt's
//     behavior of appending "%!(EXTRA ...)" -- but for this challenge,
//     simply ignore extras; you do NOT need to implement %!(EXTRA ...)).
//   - Any other verb not in the supported list should be left as-is in
//     the output, e.g. "%x" stays literally "%x" in the result, and does
//     not consume an argument.
//
// Examples:
//
//	MySprintf("%s is %d years old", "Alice", 30)
//	  -> "Alice is 30 years old"
//
//	MySprintf("Pi is %.2f", 3.14159)
//	  -> "Pi is 3.14"
//
//	MySprintf("%d%% done")
//	  -> "%!d(MISSING)% done"
//
//	MySprintf("%d", "oops")
//	  -> "%!d(string=oops)"
//
// Constraints:
//   - len(format) <= 1000
//   - 0 <= len(args) <= 20
//   - args are only ever int, string, or float64
//   - precision specifier N in %.Nf is a single ASCII digit (0-9)

// MySprintf is a simplified Sprintf supporting %d, %s, %f, %.Nf, and %%.
func MySprintf(format string, args ...interface{}) string {
	// TODO: implement
	return ""
}
