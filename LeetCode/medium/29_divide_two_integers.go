package medium

func divide(dividend int, divisor int) int {
	const MAX_INT = int(^uint(0) >> 1)
	if divisor == 0 {
		return MAX_INT
	}
	if divisor == 1 || dividend == 0 {
		return dividend
	}

	isNegative := false
	if dividend < 0 {
		isNegative = !isNegative
		dividend = -dividend
	}

	if divisor < 0 {
		isNegative = !isNegative
		divisor = -divisor
	}

	if dividend < divisor {
		return 0
	}

	result := 0
	for ; dividend > 0; result++ {
		dividend -= divisor
	}

	if dividend < 0 {
		result -= 1
	}

	if isNegative {
		result = -result
	}
	return result
}
