package main
import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, v:= range nums1 {
		m1[v] = 1
	}

	m2 := make(map[int]int)
	for _, v:= range nums2 {
		if _, ok := m1[v]; ok {
			m2[v] = 1
		}
	}

	var iSet []int
	for k :=range m2 {
		iSet = append(iSet, k)
	}

	return iSet
}

func main() {
	nums1 := [4]int {1, 2, 2, 1}
	nums2 := [2]int {2, 2}

	iSet := intersection(nums1[:], nums2[:])
	fmt.Println(iSet)
}
