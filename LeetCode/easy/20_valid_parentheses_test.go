package easy

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"[", false},
	}

	for _, test := range tests {
		got := isValid(test.in)
		if got != test.expect {
			t.Errorf("Input: %s expect: %t got %t", test.in, test.expect, got)
		}
	}
}
