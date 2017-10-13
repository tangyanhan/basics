package Easy

import (
	"testing"
	"fmt"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct{
		data []int
		expect bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, 3, 4, 4}, false},
		{[]int{1}, true},
		{[]int{1, 2}, false},
		{[]int{}, true},
		{[]int{1, 2, 3, 3, -1, 2, -1}, false},
		{[]int{1, 3, -1, 2, 2, -1, 3}, false},
		{[]int{5,4,1,-1,1,-1,4,-1, -1, 2,-1,-1, -1, 2,-1}, false},
	}

	for _, test := range tests {
		tree := CreateTreeFromSlice(test.data, 0)
		fmt.Println("==============Tree")
		PrintTree(tree, 0)
		fmt.Println("==============Tree End")
		got := isSymmetric(tree)
		if got != test.expect {
			t.Errorf("Tree: %d expect:%t got:%t", test.data, test.expect, got)
		}
	}
}
