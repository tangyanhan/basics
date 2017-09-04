package main

import "fmt"


func romanToInt(s string) int {
	NUM_MAP := map[int32]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	lastValue := 1000
	total := 0
	for _,c := range s{
		currentValue := NUM_MAP[c]
		total += currentValue
		if currentValue > lastValue {
			total -= lastValue * 2
		}
		lastValue = currentValue
	}

	return total
}

func main() {
	tests := map[string]int{"I": 1, "II": 2, "IV": 4, "V": 5, "VIII": 8, "X": 10, "IX": 9, "XIV": 14, "CD": 400}

	for roman, value := range tests {
		v := romanToInt(roman)
		if v != value {
			fmt.Println("#Error: ", roman, " Expect=", value, "Got=", v)
		}else{
			fmt.Println("PASS: ", roman, " Expect=", value, "Got=", v)
		}
	}
}