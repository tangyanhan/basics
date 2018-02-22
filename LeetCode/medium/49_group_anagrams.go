package medium

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int
		for _, c := range str {
			offset := uint32(c - 'a')
			key[offset]++
		}
		groupMap[key] = append(groupMap[key], str)
	}

	result := make([][]string, 0)
	for _, strList := range groupMap {
		result = append(result, strList)
	}
	return result
}
