package main

import "fmt"

func isNumber(s string) bool {

}

func main() {
	tests := map[string] bool {
		"0": true,
		" 0.1": true,
		"abc": false,
		"1 a": false,
		"2e10": true,
	}

	for s, expect := range tests {
		got := isNumber(s)
		if got != expect {
			fmt.Println("FAIL s=", s, "Expect=", expect, "Got=", got)
		}else{
			fmt.Println("PASS s=", s)
		}
	}
}