package main

import "fmt"
import (
	"strings"
)

type RomanBase struct {
	roman string
	base int
}

func intToRoman(num int) string {
	levels := []RomanBase{
		{"M", 1000},
		{"D", 500},
		{"C", 100},
		{"L", 50},
		{"X", 10},
		{"V", 5},
		{"I", 1},
	}

	leftMinusMap := map[int]string {
		1: "I",
		10: "X",
		100: "C",
	}

	r := ""

	for i:=0; i<len(levels) && num > 0; {
		romanBase := levels[i]
		if num >= romanBase.base {
			repeat := num/romanBase.base
			num -= repeat * romanBase.base
			// If it repeats 4 times, we have to use higher level
			if repeat == 4 && i != 0 {
				r += romanBase.roman + levels[i-1].roman
			}else{
				r += strings.Repeat(romanBase.roman, repeat)
			}
		}else{
			base := 1
			numBase := 0
			for ;;base *= 10 {
				k := num/base
				if k<10 {
					numBase = base * k
					break
				}
			}

			if numBase > 0 {
				diff := romanBase.base - numBase
				// Left minus canonly be applied when diff is 1, 10 or 100
				if roman, ok := leftMinusMap[diff]; ok {
					ratio := romanBase.base / diff
					if ratio <= 10 {
						r += roman + romanBase.roman
						num -= numBase
					}
				}
			}
			base /= 10
			i++
		}
	}

	return r
}

func main() {
	tests := map[int]string{1:"I", 2:"II", 4:"IV", 5:"V", 8:"VIII", 10:"X", 9:"IX", 14: "XIV", 400:"CD", 19:"XIX", 91:"XCI"}

	for num, roman := range tests {
		v := intToRoman(num)
		if v != roman {
			fmt.Println("#Error: ", num, " Expect=", roman, "Got=", v)
		}else{
			fmt.Println("PASS: ", num, " Expect=", roman, "Got=", v)
		}
	}
}