package problem1221

func balancedStringSplit(s string) int {
	result := 0
	len := 0

	for _, r := range s {
		if r == 'L' {
			len++
		}
		if r == 'R' {
			len--
		}
		if len == 0 {
			result++
		}
	}

	return result
}
