package Easy


func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j := 0, 0
	for j<n {
		// Find insertion point i in nums1
		for i < m && nums1[i] <= nums2[j] {
			i++
		}
		// Find range in nums2 that to be inserted
		start := j
		if i==m {  // If we should insert at end of nums1, then just give this
			j = n
		}else{
			for j<n && nums1[i] > nums2[j] {
				j++
			}
		}

		length := j-start
		// Move i:m by length
		for k:=m-1; k>=i; k-- {
			nums1[k+length] = nums1[k]
		}
		// Copy parts from nums2 to nums1
		for k:=0; k<length; k++ {
			nums1[i+k] = nums2[start + k]
		}
		i += length
		m += length
	}
}