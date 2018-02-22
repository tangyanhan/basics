package easy

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		data   []int
		expect int
	}{
		{[]int{2, 1, 1, 2}, 4},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		got := rob(test.data)
		if got != test.expect {
			t.Errorf("Data: %d  expect: %d got:%d", test.data, test.expect, got)
		}
	}
}
