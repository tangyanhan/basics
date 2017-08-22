package main

import "fmt"

func CountingSort(a []int, k int) []int {
	c := make([]int, k+1)
	b := make([]int, len(a))

	for j:=0; j<len(a); j++ {
		c[a[j]] += 1
	}

	fmt.Println("C:", c)

	for i:=1; i<=k; i++ {
		c[i] = c[i] + c[i-1]
	}

	fmt.Println("Accumulated C:", c)

	for j:=len(a)-1; j>=0; j-- {
		fmt.Println("j=", j, "a[j]=", a[j])
		fmt.Println("c[", a[j], "]=", c[a[j]])
		b[c[a[j]]-1] = a[j]
		fmt.Println("B:", b)
		c[a[j]] -= 1
	}
	return b
}

func main() {
	a := []int{2,5, 3, 0, 2, 3, 0, 3}
	b := CountingSort(a, 5)
	fmt.Println("Sorted: ", b)
}
