package medium

import (
	"log"
	"testing"
)

func TestPermute(t *testing.T) {
	l := []struct {
		in     []int
		expect [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		}},
		{[]int{1, 1, 2}, [][]int{
			[]int{1, 1, 2},
			[]int{1, 2, 1},
			[]int{2, 1, 1},
		}},
	}
	for _, item := range l {
		got := permute(item.in)
		if len(item.expect) != len(got) {
			log.Fatalf("Different length. Expect %v, Got %v", item.expect, got)
		} else {
			for i, row := range item.expect {
				for j, v := range row {
					if v != got[i][j] {
						log.Fatalf("Error near (%d, %d) Expect %v, Got %v", i, j, v, got[i][j])
					}
				}
			}
		}
	}
}
