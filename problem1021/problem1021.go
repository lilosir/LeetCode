package problem1021

func removeOuterParentheses(S string) string {
	result := ""
	open := 0

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			open++
			if open != 1 {
				result += "("
			}
		} else {
			open--
			if open != 0 {
				result += ")"
			}
		}
	}
	return result
}

// user strings Builder will use less memory!!
// func removeOuterParentheses(S string) string {
//     var level = 0
// 	var result strings.Builder
//     result.Grow(len(S))

// 	for i := 0; i < len(S); i++ {
// 		var c = S[i]
// 		if c == '(' {
// 			level++
// 			if level > 1 {
// 				result.WriteByte(c)
// 			}
// 		} else {
// 			level--
// 			if level > 0 {
// 				result.WriteByte(c)
// 			}
// 		}
// 	}
// 	return result.String()
// }
