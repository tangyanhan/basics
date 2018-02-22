package easy

import "testing"

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		x      int
		expect bool
	}{
		{0, true},
		{1, true},
		{9, true},
		{10, false},
		{11, true},
		{99, true},
		{100, false},
		{101, true},
		{19422491, true},
		{1000021, false},
		{-19422491, false},
		{-1, false},
		{-9, false},
		{-2147447412, false},
	}

	for i, test := range tests {
		got := isPalindrome(test.x)
		if got != test.expect {
			t.Errorf("#%d: x=%d expect:%t got:%t", i, test.x, test.expect, got)
		}
	}
}
