package problem657

func judgeCircle(moves string) bool {
	vertical := 0
	horizontal := 0

	for _, v := range moves {
		if v == 'U' {
			vertical++
		}
		if v == 'D' {
			vertical--
		}
		if v == 'L' {
			horizontal--
		}
		if v == 'R' {
			horizontal++
		}
	}
	return vertical == 0 && horizontal == 0
}
