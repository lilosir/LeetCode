package problem1002

func commonChars(A []string) []string {
	res := []string{}
	all := make([][]rune, len(A))
	for i := 0; i < len(A); i++ {
		all[i] = make([]rune, 26)
	}

	for i, a := range A {
		for j := 0; j < len(a); j++ {
			all[i][a[j]-97]++
		}
	}
	for i := range all[0] {
		count := all[0][i]
		if count == 0 {
			continue
		} else {
			for j := 1; j < len(all); j++ {
				if all[j][i] < count {
					count = all[j][i]
				}
			}
			if count > 0 {
				for k := 0; k < int(count); k++ {
					res = append(res, string(i+97))
				}
			}
		}

	}
	return res
}
