package medium

import "testing"

func TestGroupAnagrams(t *testing.T) {
	l := [][]string{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[]string{"bob", "bo", "oob"},
	}

	for _, strs := range l {
		result := groupAnagrams(strs)
		t.Logf("%v", result)
	}
}
