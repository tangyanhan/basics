package Easy

import "testing"

func IsEqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct{
		data []int
		expect []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 2, 3}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		input := CreateListFromSlice(test.data)
		deleteDuplicates(input)
		got := input.ToSlice()
		if !IsEqualSlice(test.expect, got) {
			t.Errorf("Input: %d Expect: %d Got: %d", test.data, test.expect, got)
		}
	}
}
