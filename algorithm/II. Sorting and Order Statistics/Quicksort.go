package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomizeData(data []int){
	rand.Seed(time.Now().Unix())

	L := len(data)
	for idx:=0; idx < L; idx++ {
		swapIdx := rand.Intn(L)
		data[idx], data[swapIdx] = data[swapIdx], data[idx]
	}
}

func generateRandomSlice(num int) []int{
	slice := make([]int, num)
	rand.Seed(time.Now().Unix())
	for i:=0; i<num; i++ {
		slice[i] = rand.Intn(num)
	}
	return slice
}

func validateSortedSlice(slice []int, isAsending bool) bool {
	if len(slice) <= 1 {
		return true
	}
	for i:=1; i<len(slice); i++ {
		if isAsending {
			if slice[i] < slice[i-1] {
				return false
			}
		}else{
			if slice[i] > slice[i-1] {
				return false
			}
		}
	}
	return true
}


// Quicksort
func partition(a []int, p int, r int) int {
	fmt.Println("#partition p=", p, "r=", r, "a=", a)
	pivot := a[r]  // Pivot value
	i := p - 1
	for j:=p; j<r; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i+1
}

func recursiveQuicksort(a []int, p int, r int){
	if p < r {
		q := partition(a, p, r)
		fmt.Println("#recursiveQuicksort: p=", p, "r=", r, "q=", q, " a=", a)
		recursiveQuicksort(a, p, q-1)
		recursiveQuicksort(a, q+1, r)
	}
}

func randomPartition(a []int, p int, r int) int {
	ri := rand.Intn(r-p) + p
	a[ri], a[r] = a[r], a[ri]
	return partition(a, p, r)
}

func randomQuicksort(a []int, p int, r int) {
	if p < r {
		q := randomPartition(a, p, r)
		randomQuicksort(a, p, q-1)
		randomQuicksort(a, q+1, r)
	}
}

func main()  {
	a := generateRandomSlice(5)
	fmt.Println("Original:", a)
	recursiveQuicksort(a, 0, len(a)-1)
	fmt.Println("Recursively sorted: ", validateSortedSlice(a, true), " Slice:", a)

	a = generateRandomSlice(5)
	randomQuicksort(a, 0, len(a)-1)
	fmt.Println("Randomly quicksorted:", validateSortedSlice(a, true), " Slice:", a)
}