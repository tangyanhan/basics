package main

// This code piece is used to solve problem 8.2-4 in MIT Introduction to Algorithms 3rd.

import "fmt"

func BuildCounting(a []int, k int) ([]int, []int) {
	c := make([]int, k+1)
	for j:=0; j<len(a); j++ {
		c[a[j]] += 1
	}

	c2 := make([]int, k+1)
	copy(c2, c)
	for i:=1; i<=k; i++ {
		c2[i] = c2[i-1] + c2[i]
	}
	return c, c2
}

func GetCounting(c []int, c2 []int, a int, b int) int{
	return c2[b] - c2[a] + c[a]
}


func main() {
	a := []int{2, 5, 3, 0, 2, 3, 0, 3}
	c, c2 := BuildCounting(a, 5)
	n := GetCounting(c, c2, 0, 5)
	fmt.Println("Count[0,5]=", n)
	n = GetCounting(c, c2, 2, 3)
	fmt.Println("Count[2,3]=", n)
}


