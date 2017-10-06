package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minStr := strs[0]
	// Find shortest string
	for i:=1; i<len(strs); i++ {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}

	for _, s := range strs {
		for i, c := range minStr {
			if uint8(c) != s[i] {
				minStr = minStr[:i]
				break
			}
		}
		if len(minStr) == 0 {
			return ""
		}
	}
	return minStr
}

func main() {
	tests := [][]string{
		{"abc", "abcd", "abcdef", "abcdefg"},
		{"aace", "akkksdf", "asdfsdf"},
		{"back", "kark"},
	}

	for i, strs := range tests {
		prefix := longestCommonPrefix(strs)
		fmt.Println("#", i, " prefix=", prefix)
	}
}
