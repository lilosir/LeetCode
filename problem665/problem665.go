package problem665

func checkPossibility(nums []int) bool {
	if isNonDecreasing(nums) {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if isNonDecreasing(nums[i+1:]) {
				if i == 0 || i == len(nums)-2 {
					return true
				}
				if nums[i] <= nums[i+2] || nums[i-1] <= nums[i+1] {
					return true
				}
				return false
			}
			return false
		}
	}

	return false
}

func isNonDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
