package main

import "testing"

func TestCompress(t *testing.T) {
	tests := []struct{
		a []byte
		b []byte
		l int
	}{
		{[]byte{'a', 'a', 'b'},[]byte{'a', '2', 'b'}, 3},
		{[]byte{'G', 't', 'W', 'Y', 'v', '&', ':', 'd', '#', 'k'},
		 []byte{'G', 't', 'W', 'Y', 'v', '&', ':', 'd', '#', 'k'}, 10},
	}

	for _, test := range tests {
		got := compress(test.a)
		if got != test.l {
			t.Fatalf("Input %d length expect %d got %d", test.a, test.l, got)
		}
		for i, n := range test.b {
			if test.b[i] != n {
				t.Fatalf("Input %d diff at index #%d: expect %d got %d", test.a, i, test.b, test.a)
			}
		}
	}
}