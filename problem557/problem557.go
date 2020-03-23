package problem557

func reverseWords(s string) string {
	b := []byte(s)
	i, j := 0, 0
	for i < len(b)-1 {
		if b[j+1] != ' ' {
			j++
		}
		if j == len(b)-1 || b[j+1] == ' ' {
			end := j
			for end > i {
				b[i], b[end] = b[end], b[i]
				i++
				end--
			}
			if j != len(b)-1 {
				j += 2
			}
			i = j
		}
	}
	return string(b)
}
