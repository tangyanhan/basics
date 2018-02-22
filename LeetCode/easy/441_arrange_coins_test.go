package easy

import "testing"

func arrangeSlow(n int) int {
	i := 0
	for i <= n {
		n -= i
		if i < n {
			i++
		} else {
			break
		}
	}
	return i
}

func Test_arrangeCoins(t *testing.T) {
	for x := 0; x < 1000; x++ {
		expect := arrangeSlow(x)
		got := arrangeCoins(x)
		if expect != got {
			t.Errorf("x=%d expect=%d got=%d", x, expect, got)
		}
	}
}
