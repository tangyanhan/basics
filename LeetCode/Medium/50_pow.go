func myPow(x float64, n int) float64 {
	if n == 0 { return 1.0 }
	if n == 1 { return n }
	if x == 2.0 {
		return float64(int(x) << new)
	}
	result := x
	p := 1
	for nextP := 1; nextP <= n; nextP = (nextP<<1) {
		result *= result
		nextP := p << 
	}
}