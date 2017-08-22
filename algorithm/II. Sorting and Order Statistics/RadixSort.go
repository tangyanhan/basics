package main

import "fmt"

// Radix sort on array a, with max element digit of d in decimal
func RadixSort(a []int, d int, base int) []int{
	b := make([]int, len(a))
	c := make([]int, base+1)

	divider := 1
	for i:=0; i<d; i++ {
		//Initiate counter
		for ic:=0; ic<len(c); ic++ {
			c[ic] = 0
		}
		//Build counter
		for j:=0; j<len(a); j++ {
			x := a[j]/divider
			radix := x%base
			c[radix] += 1
			fmt.Println("i=", i, "num=", a[j], "divider=", divider, "radix=", radix)
		}

		for ic:=1; ic<len(c); ic++ {
			c[ic] = c[ic] + c[ic-1]
		}

		for j:=len(b)-1; j>=0; j-- {
			x := a[j]/divider
			r := x%base
			b[c[r]-1] = a[j]
			c[r] -= 1
		}
		b, a = a, b
		divider *= base
	}
	return a
}

func main() {
	a := []int{853, 817, 466, 56, 341, 622, 27, 256, 281, 208, 783, 292, 834, 627, 498, 980, 724, 692, 765, 984, 261, 724, 818, 848, 546, 159, 229, 825, 704, 18}
	b := RadixSort(a, 3, 10)
	fmt.Println("Radix Sorted:", b)
}
