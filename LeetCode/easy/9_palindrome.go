package easy

func isPalindrome(x int) bool {
	tx, y := x, 0
	for tx > 0 {
		y = y*10 + (tx % 10)
		tx = tx / 10
	}
	return y == x
}
