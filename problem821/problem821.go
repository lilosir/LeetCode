package problem821

import "math"

func shortestToChar(S string, C byte) []int {
	res := make([]int, len(S))
	cpos := -1
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			cpos = i
		}
		if cpos == -1 {
			res[i] = len(S) - i - 1
		} else {
			res[i] = i - cpos
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			cpos = i
		}
		if res[i] > int(math.Abs(float64(cpos-i))) {
			res[i] = int(math.Abs(float64(cpos - i)))
		}
	}
	return res
}
