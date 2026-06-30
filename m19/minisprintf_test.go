package m19

// func TestMySprintf(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		format string
// 		args   []interface{}
// 		want   string
// 	}{
// 		{
// 			name:   "basic string and int",
// 			format: "%s is %d years old",
// 			args:   []interface{}{"Alice", 30},
// 			want:   "Alice is 30 years old",
// 		},
// 		{
// 			name:   "float with explicit precision",
// 			format: "Pi is %.2f",
// 			args:   []interface{}{3.14159},
// 			want:   "Pi is 3.14",
// 		},
// 		{
// 			name:   "literal percent plus missing arg",
// 			format: "%d%% done",
// 			args:   []interface{}{},
// 			want:   "%!d(MISSING)% done",
// 		},
// 		{
// 			name:   "type mismatch on %d",
// 			format: "%d",
// 			args:   []interface{}{"oops"},
// 			want:   "%!d(string=oops)",
// 		},
// 		{
// 			name:   "missing arg for %s",
// 			format: "Hello, %s!",
// 			args:   []interface{}{},
// 			want:   "Hello, %!s(MISSING)!",
// 		},
// 		{
// 			name:   "second of two verbs missing its arg",
// 			format: "%s and %s",
// 			args:   []interface{}{"only"},
// 			want:   "only and %!s(MISSING)",
// 		},
// 		{
// 			name:   "no verbs at all",
// 			format: "no verbs here",
// 			args:   []interface{}{},
// 			want:   "no verbs here",
// 		},
// 		{
// 			name:   "unsupported verb left literal and consumes no arg",
// 			format: "%x stays literal",
// 			args:   []interface{}{},
// 			want:   "%x stays literal",
// 		},
// 		{
// 			name:   "zero precision float rounds to nearest int with no decimal point",
// 			format: "%.0f",
// 			args:   []interface{}{3.7},
// 			want:   "4",
// 		},
// 		{
// 			name:   "default float precision is six decimal places",
// 			format: "%f",
// 			args:   []interface{}{2.5},
// 			want:   "2.500000",
// 		},
// 		{
// 			name:   "multiple verbs of different kinds in one format string",
// 			format: "%d items at %.2f each = %.2f total",
// 			args:   []interface{}{3, 1.5, 4.5},
// 			want:   "3 items at 1.50 each = 4.50 total",
// 		},
// 		{
// 			name:   "extra args beyond verb count are ignored",
// 			format: "extra args %d",
// 			args:   []interface{}{5, 99, "ignored"},
// 			want:   "extra args 5",
// 		},
// 		{
// 			name:   "empty format string with no args",
// 			format: "",
// 			args:   []interface{}{},
// 			want:   "",
// 		},
// 		{
// 			name:   "consecutive percent signs",
// 			format: "100%%",
// 			args:   []interface{}{},
// 			want:   "100%",
// 		},
// 		{
// 			name:   "type mismatch on %f given an int",
// 			format: "%f",
// 			args:   []interface{}{5},
// 			want:   "%!f(int=5)",
// 		},
// 		{
// 			name:   "type mismatch on %s given a float",
// 			format: "%s",
// 			args:   []interface{}{1.23},
// 			want:   "%!s(float64=1.23)",
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got := MySprintf(tc.format, tc.args...)
// 			if got != tc.want {
// 				t.Errorf("MySprintf(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
// 			}
// 		})
// 	}
// }
