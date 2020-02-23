package problem1266

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	for i := 0; i < len(points)-1; i++ {
		x := points[i][0] - points[i+1][0]
		y := points[i][1] - points[i+1][1]
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}

		distance := 0
		if x > y {
			distance = x
		} else {
			distance = y
		}

		result += distance
	}

	return result
}
