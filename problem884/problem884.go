package problem884

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	words := make(map[string]int)
	aSli := strings.Split(A, " ")
	bSli := strings.Split(B, " ")
	for _, s := range aSli {
		words[s]++
	}
	for _, s := range bSli {
		words[s]++
	}

	res := []string{}
	for k, v := range words {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
