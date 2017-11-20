package Easy


func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	solution := make([]int, len(nums))
	for i, v := range nums {
		// Try to pick value at i, so we should add to non-adjectie solution at i-2
		nonAdjSolution := 0  // Solution at i-2
		if i > 1 {
			nonAdjSolution = solution[i-2]
		}
		pickI := nonAdjSolution + v
		adjSolution := 0
		if i > 0 {
			adjSolution = solution[i-1]
		}
		if pickI > adjSolution {
			solution[i] = pickI
		}else{
			solution[i] = adjSolution
		}
	}

	return solution[len(nums)-1]
}