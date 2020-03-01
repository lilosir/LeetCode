package problem1309

func freqAlphabets(s string) string {
	result := ""
	i := 0
	for i < len(s) {
		if i == len(s)-2 {
			return result + string(s[i]+48) + string(s[i+1]+48)
		}
		if i == len(s)-1 {
			return result + string(s[i]+48)
		}
		if s[i+2] != '#' {
			result += string(s[i] + 48)
			i++
		} else {
			result = result + string((s[i]-48)*10+(s[i+1]-48)+96)
			i += 3
		}
	}
	return result
}
