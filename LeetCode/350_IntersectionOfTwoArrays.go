package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	var iSet[]int
	i, j := 0, 0
	for i<len(nums1) && j<len(nums2) {
		if nums1[i] > nums2[j] {
			j++
			continue
		}else if nums1[i] < nums2[j] {
			i++
			continue
		}else{
			iSet = append(iSet, nums1[i])
			i++
			j++
		}
	}
	return iSet
}

func main() {
	nums1 := [4]int {1, 2, 2, 1}
	nums2 := [2]int {2, 2}

	iSet := intersect(nums1[:], nums2[:])
	fmt.Println(iSet)
}