import testing
import math

func TestMyPow(t *testing.T) {
	testCases := []struct{
		x float64
		n int
		expect float64
	}{
		{0.0, 0, 1.0},
		{0, 1, 0.0},
		{2.0, 10, 1024.0},
		{2.10000, 3, 9.26100},
	}

	for _, testCase := range testCases {
		got := myPow(testCase.x, testCase.n)
		if math.abs(got-expect) > 0.00001 {
			t.Fatalf("x=%f n=%d expect=%f got=%f", testCase.x, testCase.n, testCase.expect, testCase.got)
		}
	}
}