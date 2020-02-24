package problem709

func toLowerCase(str string) string {
	asciiStr := []byte(str)
	for i, v := range str {
		if v >= 65 && v <= 90 {
			asciiStr[i] = byte(v + 32)
		}
	}
	return string(asciiStr)
}
