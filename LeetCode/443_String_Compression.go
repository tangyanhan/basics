package main

import "strconv"

func compress(chars []byte) int {
	i := 0
	k := 0
	for i<len(chars) {
		chars[k] = chars[i]
		k++
		j := i
		i++
		for i<len(chars) && chars[i] == chars[j] {
			i++
		}
		dist := i - j
		if dist > 1 {
			s := strconv.Itoa(dist)
			for l, n := range s {
				chars[k+l] = byte(n)
			}
			k += len(s)
		}
	}
    return k
}