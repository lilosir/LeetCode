package problem1200

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	absDif := math.Abs(float64(arr[0] - arr[1]))
	res := [][]int{{arr[0], arr[1]}}
	for i := 1; i < len(arr)-1; i++ {
		if math.Abs(float64(arr[i]-arr[i+1])) < absDif {
			absDif = math.Abs(float64(arr[i] - arr[i+1]))
			res = [][]int{{arr[i], arr[i+1]}}
		} else if math.Abs(float64(arr[i]-arr[i+1])) == absDif {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}
