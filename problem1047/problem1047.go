package problem1047

import "strings"

func removeDuplicates(S string) string {
	strSlice := []string{}
	for i := 0; i < len(S); i++ {
		if len(strSlice) == 0 || strSlice[len(strSlice)-1] != string(S[i]) {
			strSlice = append(strSlice, string(S[i]))
		} else {
			strSlice = strSlice[:len(strSlice)-1]
		}
	}
	return strings.Join(strSlice, "")
}
