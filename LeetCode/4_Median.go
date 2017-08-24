package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0.0
	}
	medianIdx := (totalLen-1)/2

	median := 0.0
	for i,j,k:=0,0,0; i<len(nums1) || j<len(nums2); {
		hasLeft := i<len(nums1)
		hasRight := j<len(nums2)
		num := 0
		if hasLeft && hasRight {
			if nums1[i] < nums2[j] {
				num = nums1[i]
				i++
			}else{
				num = nums2[j]
				j++
			}
		}else if hasLeft{
			num = nums1[i]
			i++
		}else if hasRight{
			num = nums2[j]
			j++
		}
		if k == medianIdx {
			median = float64(num)
			if totalLen % 2 == 1 {
				break
			}
		}else if k>medianIdx {
			median = (median + float64(num))/2.0
			break
		}
		k++
	}

	return median
}

func findMedian(a []int, b []int, k int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	if len(a) == 0 {
		fmt.Println("k=", k)
		return float64(b[k-1])
	}
	if k == 1 {
		fmt.Println("a=", a, "b=", b)
		return math.Min(float64(a[0]), float64(b[0]))
	}

	i := int(math.Min(float64(len(a)), float64(k/2)))
	j := int(math.Min(float64(len(b)), float64(k/2)))

	if a[i-1] > b[j-1] {
		return findMedian(a, b[j:], k-j)
	}else{
		return findMedian(a[i:], b, k-i)
	}
	return 0.0
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	left := (totalLen+1)>>1
	right := (totalLen+2)>>2
	fmt.Println("left=", left, "right=", right)
	return (findMedian(nums1, nums2, left) + findMedian(nums1, nums2, right))/2.0
}

func main() {
	a := []int {1, 2}
	b := []int {3, 4}
	m := findMedianSortedArrays(a, b)
	fmt.Println("m=", m)
	m = findMedianSortedArrays2(a, b)
	fmt.Println("Solution 2: m=", m)
}
