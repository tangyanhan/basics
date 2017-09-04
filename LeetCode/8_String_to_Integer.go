package main

import "fmt"

func myAtoi(str string) int {
	start := -1
	end := -1
	isNegative := false
	for i,c := range str {
		if start == -1 {
			if c == '-' && (i+1) < len(str) {
				start = i+1
				isNegative = true
			}else if c == '+' && (i+1) < len(str) {
				start = i+1
			}else if c >='0' && c <='9' {
				start = i
			}else if c != ' '{
				return 0
			}
		}else if end == -1 && (c>'9' || c<'0') {
			end = i
			break
		}
	}
	if end == -1 {
		end = len(str)
	}

	if start == -1 {
		return 0
	}

	numStr := str[start: end]
	//fmt.Println("str=", str, "numStr=", numStr)
	value := 0
	for i, base :=len(numStr)-1, 1; i>=0; i, base = i-1, base * 10 {
		value += int(numStr[i] - '0') * base
		if !isNegative && value > 2147483647 {
			return 2147483647
		}else if isNegative && value > 2147483648{
			return -2147483648
		}
	}
	if isNegative{
		value = -value
	}
	return value
}

func main() {
	tests := map[string]int {
		"1234567" : 1234567,
		"-1234567": -1234567,
		"+123": 123,
		"+-2": 0,
		"aa12345bb12": 0,
		"  3345": 3345,
		"+abc3345+-": 0,
		"  -3345-+": -3345,
		"00345": 345,
		"1234444444444444444444444444": 2147483647,
		"-98797987897972893792742834": -2147483648,
	}

	for s,i := range tests {
		got := myAtoi(s)
		if got != i {
			fmt.Println("FAIL: s=", s, "expect=", i, "got=", got)
		}else{
			fmt.Println("PASS: s=", s)
		}
	}
}