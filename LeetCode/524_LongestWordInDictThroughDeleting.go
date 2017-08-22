package main
import "fmt"
import "encoding/json"
import "io/ioutil"

func findLongestWord(s string, d []string) string {
	dict := sortStringByLength(d)
	for i:=len(dict)-1; i>=0; i-- {
		if len(dict[i]) > len(s) {
			break
		}

		j := 0
		k := 0
		for ; j<len(dict[i]); j++ {
			for ;k<len(s); k++ {
				if s[k] == dict[i][j] {
					break
				}
			}
			if k==len(s) {
				break
			}
		}
		if j==len(dict[i]) {
			return dict[i]
		}
	}
	return ""
}

func sortStringByLength(d []string) []string{
	b := make([]string, len(d))
	c := make([]int, 1001)

	for i:=0; i<len(d); i++ {
		c[len(d[i])] += 1
	}

	for k:=1; k<1001; k++ {
		c[k] = c[k-1] + c[k]
	}

	for i:=0; i<len(d); i++ {
		b[c[len(d[i])] - 1] = d[i]
		c[len(d[i])] -= 1
	}
	return b
}

func readJsonFile(filename string) ([][]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	result := make([][]string, 0, 1000)
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return result, nil
}

func main() {
	arr, err := readJsonFile("LeetCode/524_test_data.json")
	if err != nil {
		panic("Failed to read json file")
	}

	for i:=0; i<len(arr); i++ {
		s := arr[i][0]
		d := arr[i][1:]
		fmt.Println("result=", findLongestWord(s, d))
	}
}