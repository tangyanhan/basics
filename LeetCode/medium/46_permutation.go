package medium

import (
	"sort"
)

func permute(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return [][]int{}
	}

	resultCap := 1
	for i := 2; i <= len(nums); i++ {
		resultCap *= i
	}
	all := make([][]int, 0, resultCap)

	nextPermuteFunc := func() bool {
		var i int
		for i = len(nums) - 2; (i >= 0) && (nums[i] >= nums[i+1]); i-- {

		}
		if i < 0 {
			return false
		}

		var k int
		for k = len(nums) - 1; (k > i) && (nums[k] <= nums[i]); k-- {

		}
		nums[i], nums[k] = nums[k], nums[i]
		remain := nums[i+1:]
		for j := 0; j < (len(remain) / 2); j++ {
			k := len(remain) - j - 1
			remain[j], remain[k] = remain[k], remain[j]
		}
		return true
	}

	cpy := make([]int, len(nums))
	copy(cpy, nums)
	all = append(all, cpy)
	for nextPermuteFunc() {
		cpy = make([]int, len(nums))
		copy(cpy, nums)
		all = append(all, cpy)
	}
	return all
}
