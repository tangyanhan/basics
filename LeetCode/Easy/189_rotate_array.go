package Easy


func rotate(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}

	k %= len(nums)
	if k == 0 {
		return
	}

	pivot := len(nums) - k
	reverse(nums, 0, pivot-1)
	reverse(nums, pivot, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}


func reverse(a []int, start, end int) {
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}