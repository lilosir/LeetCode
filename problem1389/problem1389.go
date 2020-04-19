package problem1389

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

	for i := range index {
		if res[index[i]] != -1 {
			moveBack(index[i], &res)
		}
		res[index[i]] = nums[i]
	}
	return res
}

func moveBack(i int, s *[]int) {
	buf, count := 0, 1
	for j := i + 1; j < len(*s); j++ {
		if count == 0 {
			break
		}
		if (*s)[j] != -1 {
			count++
		} else {
			count--
		}
		buf++
	}
	for k := i + buf; k > i; k-- {
		(*s)[k] = (*s)[k-1]
	}
}
