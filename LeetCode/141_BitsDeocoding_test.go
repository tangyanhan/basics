package main

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	tests := []struct{
		a []int
		expect bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
		{[]int{1, 0, 1, 1, 0}, true},
	}

	for _, test := range tests {
		got := isOneBitCharacter(test.a)
		if got != test.expect {
			t.Fatalf("Input %d expect %t got %t", test.a, test.expect, got)
		}
	}
}