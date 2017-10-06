package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i:=0; i<len(nums)-2; i++ {
		if i>0 && nums[i-1] == nums[i] {  // Same integer, skip and look for the next
			continue
		}
		// Search other two integers from two directions in remaining parts
		j := i+1
		k := len(nums)-1
		need := 0 - nums[i]
		for j < k {
			sum := nums[j] + nums[k]

			if sum == need{
				result = append(result, []int{nums[i], nums[j], nums[k]})
				// Skip other same combinations
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			}else if sum > need {  // The larger number is way too high, make it smaller
				k--
			}else{  // sum < 0, we need to make it larger
				j++
			}
		}
	}

	return result
}


func readJsonFile(filename string) ([][]int, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	result := make([][]int, 0, 1000)
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return result, nil
}

func main() {
	arr, err := readJsonFile("15_data.txt")
	if err != nil {
		panic("Failed to read json file")
	}

	for _, nums := range arr {
		start := time.Now()
		result := threeSum(nums)
		end := time.Now()
		elapsed := end.Sub(start)
		//for j, set := range result {
		//	fmt.Println("#", i, "##", j, set)
		//}
		fmt.Println("Elapsed: ", elapsed)
		fmt.Println("Result len:", len(result), "Input len:", len(nums))
	}
}

