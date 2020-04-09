package problem1047

func removeDuplicates(S string) string {
	strSlice := []byte{}
	for i := 0; i < len(S); i++ {
		if len(strSlice) == 0 || strSlice[len(strSlice)-1] != S[i] {
			strSlice = append(strSlice, S[i])
		} else {
			strSlice = strSlice[:len(strSlice)-1]
		}
	}
	return string(strSlice)
}
