package problem804

import "strings"

func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool)

	for _, v := range words {
		var current strings.Builder
		for _, r := range v {
			code := morseCode[r-97]
			current.WriteString(code)
		}
		if _, ok := m[current.String()]; !ok {
			m[current.String()] = true
		}
	}

	return len(m)
}
