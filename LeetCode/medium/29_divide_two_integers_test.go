package medium

import (
	"flag"
	"testing"
)

var Param string

func init() {
	flag.StringVar(&Param, "param", "", "External value")
	flag.Parse()
}

func TestDivide(t *testing.T) {

	testCases := []struct {
		a      int
		b      int
		expect int
	}{
		{0, 0, 2147483647},
		{0, 1, 0},
		{10, 5, 2},
		{3, 3, 1},
		{1, -1, -1},
		{1, 2, 0},
		{5, 2, 2},
		{-2147483648, -1, 2147483647},
		{-8, 4, -2},
	}

	t.Fatalf("Param:%s", Param)
	for _, c := range testCases {
		got := divide(c.a, c.b)
		if c.expect != got {
			t.Fatalf("%d/%d=%d, got:%d", c.a, c.b, c.expect, got)
		}
	}
}
