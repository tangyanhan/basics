package easy

import "testing"

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		greed   []int
		cookies []int
		expect  int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	for _, test := range tests {
		got := findContentChildren(test.greed, test.cookies)
		if test.expect != got {
			t.Errorf("Greed: %d Cookies: %d expect:%d got:%d", test.greed, test.cookies, test.expect, got)
		}
	}
}
