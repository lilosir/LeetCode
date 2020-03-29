package problem1160

func countCharacters(words []string, chars string) int {
	sliceB := make([]int, 26)
	for _, v := range chars {
		sliceB[v-'a']++
	}

	res := 0
	for _, word := range words {
		if contains(word, sliceB) {
			res += len(word)
		}
	}

	return res
}

func contains(a string, b []int) bool {
	temp := make([]int, len(b))
	copy(temp, b)

	for _, v := range a {
		if temp[v-'a'] > 0 {
			temp[v-'a']--
		} else {
			return false
		}
	}
	return true
}
