package problem824

import (
	"strings"
)

func toGoatLatin(S string) string {
	strSli := strings.Split(S, " ")
	for i, s := range strSli {
		if startWithVowel(s) {
			s += "ma"
		} else {
			s = s[1:] + string(s[0]) + "ma"
		}
		s += strings.Repeat("a", (i+1)*1)
		strSli[i] = s
	}
	return strings.Join(strSli, " ")
}

func startWithVowel(s string) bool {
	switch s[:1] {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	default:
		return false
	}
}
