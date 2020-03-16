package problem771

func numJewelsInStones(J string, S string) int {
	result := 0
	JMap := make(map[rune]int)
	for _, r := range J {
		JMap[r] = 0
	}

	for _, r := range S {
		if _, ok := JMap[r]; ok {
			result++
		}
	}

	return result
}
