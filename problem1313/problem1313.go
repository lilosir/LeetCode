package problem1313

func decompressRLElist(nums []int) []int {
	result := []int{}

	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}
