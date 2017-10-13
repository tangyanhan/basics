package Easy

import "testing"

func Test_merge(t *testing.T) {
	tests := []struct{
		nums1 []int
		nums2 []int
		expect []int
	}{
		{[]int{1, 4, 5}, []int{2, 3, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{}, []int{}, []int{}},
	}

	for _, test := range tests {
		nums1 := make([]int, len(test.nums1) + len(test.nums2))
		copy(nums1, test.nums1)
		merge(nums1, len(test.nums1), test.nums2, len(test.nums2))
		for i, v := range nums1 {
			if test.expect[i] != v {
				t.Errorf("s1: %d s2: %d expect: %d got: %d", test.nums1, test.nums2, test.expect, nums1)
				break
			}
		}
	}
}
