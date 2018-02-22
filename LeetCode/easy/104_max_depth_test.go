package easy

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		data   []int
		expect int
	}{
		{[]int{0, 0, 0, -1, -1, -1, 0, -1, -1, -1, -1, -1, -1, -1, 0}, 4},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
	}

	for _, test := range tests {
		tree := CreateTreeFromSlice(test.data, 0)
		got := maxDepth(tree)
		fmt.Println("============ Tree Begin")
		PrintTree(tree, 0)
		fmt.Println("============ Tree End")
		if got != test.expect {

			t.Errorf("expect: %d got: %d", test.expect, got)
		}
	}
}
