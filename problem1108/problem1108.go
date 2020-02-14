package problem1108

func defangIPaddr(address string) string {
	result := ""
	for _, s := range address {
		if s == '.' {
			result += "[.]"
		} else {
			result += string(s)
		}
	}
	return result
}
