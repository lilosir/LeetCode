package problem1021

import "strings"

// func removeOuterParentheses(S string) string {
// 	result := ""
// 	open := 0

// 	for i := 0; i < len(S); i++ {
// 		if S[i] == '(' {
// 			open++
// 			if open != 1 {
// 				result += "("
// 			}
// 		} else {
// 			open--
// 			if open != 0 {
// 				result += ")"
// 			}
// 		}
// 	}
// 	return result
// }

// user strings Builder will use less memory!!
func removeOuterParentheses(S string) string {
	var result strings.Builder
	open := 0

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			open++
			if open != 1 {
				result.WriteRune('(')
			}
		} else {
			open--
			if open != 0 {
				result.WriteRune(')')
			}
		}
	}
	return result.String()
}
