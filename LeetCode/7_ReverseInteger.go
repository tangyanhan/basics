package main

import "fmt"

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		x = -x
		isNegative = true
	}
	s := make([]int, 20)
	i := 0
	for ;x>0; x/=10{
		n := x % 10
		s[i] = n
		i++
	}

	var r int32 = 0
	base := 1
	for j:=i-1; j>=0; j-- {
		current := base * s[j]
		if int(int32(current)) < current {
			return 0
		}
		if int32(int(r) + int(current)) < r {
			return 0
		}
		r += int32(current)
		base *= 10
	}
	if isNegative {
		r = -r
	}
	return int(r)
}

func main() {
	tests := map[int]int {
		123: 321,
		-123: -321,
		1000000003: 0,
	}
	for x, r := range tests {
		v := reverse(x)
		if v != r {
			fmt.Println("FAIL: x=", x, "Expect=", r, "Got=", v)
		}else{
			fmt.Println("PASS: x=", x)
		}
	}
}
