package Easy

import (
	"testing"
	"fmt"
)


func isEqualSlice(a, b[]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	tests := [][]int{
		{3,   1, 2, 3, 4, 5, 6, 7, 5, 6, 7, 1, 2, 3, 4},
		{7,   1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7},
		{8,   1, 2, 3, 4, 5, 6, 7, 7, 1, 2, 3, 4, 5, 6},
		{10},
	}

	for _, test := range tests {
		k := test[0]
		n := int((len(test)-1)/2)
		nums := test[1:n+1]
		expect := test[n+1:]
		fmt.Println("expect=", expect, " input=", nums)
		rotate(nums, k)
		if !isEqualSlice(nums, expect) {
			t.Errorf("k=%d Expect %s, got %s", k, expect, nums)
		}
	}
}
