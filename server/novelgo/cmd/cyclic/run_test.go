package main

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	filename := "1.txt"
	run(&filename)
}

func TestParseLine(t *testing.T) {
	type parseLineResult struct {
		a   int
		b   int
		err error
	}

	tests := []struct {
		name  string
		input string
		want  parseLineResult
	}{
		{
			name:  "Valid input",
			input: "10 20",
			want:  parseLineResult{10, 20, nil},
		},
		{
			name:  "Invalid input (non-integer values)",
			input: "abc def",
			want:  parseLineResult{-1, -1, errors.New("")},
		},
		{
			name:  "Invalid input (not enough values)",
			input: "10",
			want:  parseLineResult{-1, -1, errors.New("")},
		},
		{
			name:  "Invalid input (too many values)",
			input: "10 20 30",
			want:  parseLineResult{-1, -1, errors.New("")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a, b, err := parseLine(test.input)
			if a != test.want.a || b != test.want.b || (err != nil && test.want.err == nil) || (err == nil && test.want.err != nil) {
				t.Errorf("Want %v, got %v", test.want, parseLineResult{a, b, err})
			}
		})
	}
}

// Replace "parseIntegers" with your actual function name.
// }

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
// func TestHelloName(t *testing.T) {
// 	name := "Gladys"
// 	want := regexp.MustCompile(`\b` + name + `\b`)
// 	msg, err := Hello("Gladys")
// 	if !want.MatchString(msg) || err != nil {
// 		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
// 	}
// }
//
// // TestHelloEmpty calls greetings.Hello with an empty string,
// // checking for an error.
// func TestHelloEmpty(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}
// }
