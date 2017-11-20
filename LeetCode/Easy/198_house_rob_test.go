package Easy


import "testing"

func TestRob(test *testing.T) {
	ts := []struct{
		in []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4}, 6},
		{[]int{5, 1, 1, 4, 3}, 9},
	}

	for _, t := range ts {
		got := rob(t.in)
		if got != t.expect {
			test.Fatalf("Input: %d expect: %d got %d\n", t.in, t.expect, got)
		}
	}
}